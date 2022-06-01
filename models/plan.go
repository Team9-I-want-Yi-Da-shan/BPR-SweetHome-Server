package models

import (
	"errors"

	"gin-Home-server/db"
	"gin-Home-server/forms"
)

type PlanModel struct{}

type Plan_person struct {
	ID          int64  `db:"id, primarykey, autoincrement" json:"plan_id"`
	UserID      int64  `db:"user_id" json:"user_id"`
	Name        string `db:"name" json:"title"`
	Description string `db:"description" json:"description"`
	Comment     string `db:"comment" json:"comment"`
}
type Plan_family struct {
	ID          int64  `db:"id, primarykey, autoincrement" json:"plan_id"`
	FamilyID    int64  `db:"family_id" json:"family_id"`
	Name        string `db:"name" json:"title"`
	Description string `db:"description" json:"description"`
	Comment     string `db:"comment" json:"comment"`
}

//Create ...
func (m PlanModel) CreatePersonalPlan(form forms.PersonalPlanForm) (planId int64, err error) {
	err = db.GetDB().QueryRow("INSERT INTO public.plan_person(user_id, name, description,comment) VALUES($1, $2, $3,$4) RETURNING id", form.PersonId, form.Name, form.Description, form.Comment).Scan(&planId)

	if err != nil {
		return 0, err
	}
	return planId, err
}

//Create ...
func (m PlanModel) CreateFamilylPlan(form forms.FamilyPlanForm) (planId int64, err error) {
	err = db.GetDB().QueryRow("INSERT INTO public.plan_family(family_id, name, description,comment) VALUES($1, $2, $3,$4) RETURNING id", form.FamilyId, form.Name, form.Description, form.Comment).Scan(&planId)

	if err != nil {
		return 0, err
	}
	return planId, err
}

//All by id personal
func (m PlanModel) AllPersonPlan(userID int64) (plans []DataList, err error) {
	_, err = db.GetDB().Select(&plans, "SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data, (SELECT row_to_json(n) FROM ( SELECT count(a.id) AS total FROM public.plan_person AS a WHERE a.user_id=$1 LIMIT 1 ) n ) AS meta FROM ( SELECT a.id, a.user_id,a.name as title, a.description, a.comment FROM public.plan_person a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 ORDER BY a.id DESC) d", userID)
	return plans, err
}

//All by id family
func (m PlanModel) AllFamilyPlan(familyID int64) (plans []DataList, err error) {
	_, err = db.GetDB().Select(&plans, "SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data, (SELECT row_to_json(n) FROM ( SELECT count(a.id) AS total FROM public.plan_family AS a WHERE a.user_id=$1 LIMIT 1 ) n ) AS meta FROM ( SELECT a.id, a.user_id,a.name as title, a.description, a.comment FROM public.plan_family a LEFT JOIN public.family u ON a.family_id = u.id WHERE a.user_id=$1 ORDER BY a.id DESC) d", familyID)
	return plans, err
}

//One personal
func (m PlanModel) GetPlanByID(id int64) (plan Plan_person, err error) {
	err = db.GetDB().SelectOne(&plan, "SELECT a.id, a.name as title, a.description, a.comment FROM public.plan_person a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.id=$1  LIMIT 1", id)
	return plan, err
}

//One family
func (m PlanModel) GetFamilyPlanByID(id int64) (plan Plan_person, err error) {
	err = db.GetDB().SelectOne(&plan, "SELECT a.id, a.name as title, a.description, a.comment FROM public.plan_family a LEFT JOIN public.family u ON a.family_id = u.id WHERE a.id=$1  LIMIT 1", id)
	return plan, err
}

//Delete ...
func (m PlanModel) DeletePersonalPlan(dbName string, id int64) (err error) {

	operation, err := db.GetDB().Exec("DELETE FROM public.plan_person WHERE id=$1", id)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("no records were deleted")
	}

	return err
}

//Delete ...
func (m PlanModel) DeleteFamilyPlan(id int64) (err error) {

	operation, err := db.GetDB().Exec("DELETE FROM public.plan_family WHERE id=$1", id)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("no records were deleted")
	}

	return err
}

//Update
func (m PlanModel) UpdatePersonPlan(id int64, form forms.PersonalPlanForm) (err error) {

	operation, err := db.GetDB().Exec("UPDATE public.plan_person SET name=$2, description=$3 , comment = $4 WHERE id=$1", id, form.Name, form.Description, form.Comment)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("updated 0 records")
	}

	return err
}

//Update family plan by plan id
func (m PlanModel) UpdateFamilyPlan(id int64, form forms.FamilyPlanForm) (err error) {

	operation, err := db.GetDB().Exec("UPDATE public.plan_family SET name=$2, description=$3 , comment = $4 WHERE id=$1", id, form.Name, form.Description, form.Comment)
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("updated 0 records")
	}

	return err
}
