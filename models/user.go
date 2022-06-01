package models

import (
	"errors"
	"log"

	"gin-Home-server/db"
	"gin-Home-server/forms"

	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"user_id"`
	FamilyID  int64  `db:"familyid" json:"family_id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"-"`
	Name      string `db:"name" json:"name"`
	UpdatedAt int64  `db:"updated_at" json:"-"`
	CreatedAt int64  `db:"created_at" json:"-"`
}

//Family
type Family struct {
	ID       int64  `db:"id, primarykey, autoincrement" json:"family_id"`
	Name     string `db:"name" json:"name"`
	Admin_id int64  `db:"admin_id" json:"admin_id"`
}

//UserModel ...
type UserModel struct{}

var authModel = new(AuthModel)

//Login ...
func (m UserModel) Login(form forms.LoginForm) (user User, token Token, err error) {

	err = db.GetDB().SelectOne(&user, "SELECT id, familyid, email, password, name, updated_at, created_at FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)

	if err != nil {
		return user, token, err
	}

	//Compare the password form and database if match
	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		log.Printf("login bcrypt err")
		return user, token, err
	}

	//Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(user.ID)
	if err != nil {
		log.Printf("login token err")
		return user, token, err
	}

	saveErr := authModel.CreateAuth(user.ID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}
	log.Printf("token", tokenDetails.AccessToken)
	log.Printf("token", tokenDetails.RefreshToken)

	return user, token, nil
}

//Register ...
func (m UserModel) Register(form forms.RegisterForm) (user User, err error) {
	getDb := db.GetDB()

	//Check if the user exists in database
	checkUser, err := getDb.SelectInt("SELECT count(id) FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	if checkUser > 0 {
		return user, errors.New("email already exists")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	//Create the user and return back the user ID
	err = getDb.QueryRow("INSERT INTO public.user(email, password, name) VALUES($1, $2, $3) RETURNING id", form.Email, string(hashedPassword), form.Name).Scan(&user.ID)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	user.Name = form.Name
	user.Email = form.Email

	return user, err
}

//One ...
func (m UserModel) One(userID int64) (user User, err error) {
	err = db.GetDB().SelectOne(&user, "SELECT id,familyId, email, name FROM public.user WHERE id=$1 LIMIT 1", userID)
	return user, err
}

func (m UserModel) OneFamily(familyID int64) (family Family, err error) {
	err = db.GetDB().SelectOne(&family, "SELECT id AS family_id,name, admin_id FROM public.family WHERE id=$1 LIMIT 1", familyID)
	return family, err
}

//AllFamily
func (m UserModel) GetFamilyByName(name string) (families []DataList, err error) {
	_, err = db.GetDB().Select(&families, "SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data,(SELECT row_to_json(n) FROM ( SELECT count(a.id) AS total FROM public.family AS a WHERE  position($1 in a.name)>0 LIMIT 1 ) n ) AS meta FROM ( SELECT a.id AS family_id,a.name, json_build_object('user_id', u.id, 'name', u.name, 'email', u.email) AS admin FROM public.family a LEFT JOIN public.user u ON a.admin_id = u.id WHERE  position($1 in a.name)>0 ORDER BY a.id DESC) d", name)
	return families, err
}

//Search Family
func (m UserModel) getFamilyByFamilyID(id int64) (family Family, err error) {
	err = db.GetDB().SelectOne(&family, "SELECT (row_to_json(d)) AS data FROM ( SELECT a.id AS family_id,a.name, json_build_object('user_id', u.id, 'name', u.name, 'email', u.email) AS admin FROM public.family a LEFT JOIN public.user u ON a.admin_id = u.id  WHERE a.id = $1)d", id)
	return family, err
}

//Add Family
func (m UserModel) CreateFamily(form forms.FamilyForm) (family Family, err error) {
	tx, err := db.GetDB().Begin()
	if err != nil {
		return family, err
	}
	defer tx.Commit()

	err = tx.QueryRow("INSERT INTO public.family(name, admin_id) VALUES($1, $2) RETURNING id", form.Name, form.Admin_id).Scan(&family.ID)
	if err != nil {
		tx.Rollback()
		return family, errors.New("something went wrong, please try again later")
	}
	operation, err := tx.Exec("UPDATE public.user SET  familyid = $2  WHERE id=$1 ", form.Admin_id, family.ID)
	success, _ := operation.RowsAffected()
	if success == 0 || err != nil {
		tx.Rollback()
		return family, errors.New("failed to update admin status, please try again later")
	}

	family.Name = form.Name
	family.Admin_id = int64(form.Admin_id)

	return family, err
}

//AddFamilyUser
func (m UserModel) AddFamilyUser(form forms.RequestAddFamilyForm) (err error) {
	checkUser, err := db.GetDB().SelectInt("SELECT count(id) FROM public.response WHERE admin_id=$1 AND request_id=$2 AND finish =0 LIMIT 1", form.Admin_id, form.User_id)
	if err != nil {
		return errors.New("something went wrong, please try again later")
	}

	if checkUser > 0 {
		return errors.New("response already exists")
	}
	err = db.GetDB().QueryRow("INSERT INTO public.response(admin_id,request_id,confirm, finish) VALUES($1, $2,0,0) ", form.Admin_id, form.User_id).Err()
	return err
}

//AllAdminNewResponse
func (m UserModel) AllAdminNewResponse(userID int64) (responses []HomeList, err error) {
	_, err = db.GetDB().Select(&responses, "SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data FROM ( SELECT a.id AS family_id ,a.admin_id,a.request_id, a.updated_at, json_build_object('user_id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.response a LEFT JOIN public.user u ON a.request_id = u.id WHERE a.admin_id=$1 AND a.finish=0 ORDER BY a.updated_at DESC) d", userID)
	return responses, err
}

//ConfirmUserToFamily
func (m UserModel) ConfirmUserToFamily(form forms.ConfirmRequestAddFamilyForm) (err error) {
	//change state of response,update
	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	operation, err := tx.Exec("UPDATE public.response SET  confirm = $3, finish = 1 WHERE admin_id=$1 AND request_id=$2", form.Admin_id, form.User_id, form.Confirm)
	success, _ := operation.RowsAffected()
	if success == 0 || err != nil {
		tx.Rollback()
		return errors.New("updated 0 response records")
	}
	if form.Confirm == 0 {
		return nil
	}
	//update user family id
	ope, err := tx.Exec("UPDATE public.user SET  familyid = $2  WHERE  id=$1", form.User_id, form.Family_id)
	successCnt, _ := ope.RowsAffected()
	if successCnt == 0 || err != nil {
		tx.Rollback()
		return errors.New("updated 0 user records")
	}
	return err
}

//familyListByFamilyID
func (m UserModel) FamilyListByFamilyID(familyID int64) (responses []HomeList, err error) {
	_, err = db.GetDB().Select(&responses, "SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data FROM ( SELECT a.id AS user_id ,a.name,a.email FROM public.user a   WHERE a.familyid=$1  ORDER BY a.updated_at DESC) d", familyID)
	return responses, err
}
