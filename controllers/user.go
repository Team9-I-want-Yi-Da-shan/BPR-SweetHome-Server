package controllers

import (
	"gin-Home-server/forms"
	"gin-Home-server/models"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

//UserController ...
type UserController struct{}

var userModel = new(models.UserModel)
var userForm = new(forms.UserForm)

//getUserID ...
func getUserID(c *gin.Context) (userID int64) {
	//MustGet returns the value for the given key if it exists, otherwise it panics.
	return c.MustGet("userID").(int64)
}

//Login ...
func (ctrl UserController) Login(c *gin.Context) {
	var loginForm forms.LoginForm

	if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
		message := userForm.Login(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, token, err := userModel.Login(loginForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid login details", "user": user, "err": err})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": user, "token": token})
}

//Register ...
func (ctrl UserController) Register(c *gin.Context) {
	var registerForm forms.RegisterForm

	if validationErr := c.ShouldBindJSON(&registerForm); validationErr != nil {

		message := userForm.Register(validationErr)

		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})

		return
	}

	user, err := userModel.Register(registerForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered", "user": user})
}

//Logout ...
func (ctrl UserController) Logout(c *gin.Context) {

	au, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User not logged in"})
		return
	}

	deleted, delErr := authModel.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

//getUserByID
func (ctrl UserController) SearchUserByID(c *gin.Context) {
	id := c.Param("id")
	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		return
	}
	result, err := userModel.One(getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Could not get user", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": result})
}

//FindFamily by name
func (ctrl UserController) SearchFamilyByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid search parameter", "name": name})
		return
	}
	results, err := userModel.GetFamilyByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Could not get family", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

//Find family by id
func (ctrl UserController) SearchFamilyById(c *gin.Context) {
	id := c.Param("id")
	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		return
	}
	result, err := userModel.OneFamily(getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Could not get family", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"family": result})
}

// CreateFamily admin
func (ctrl UserController) CreateFamily(c *gin.Context) {
	var form forms.FamilyForm
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": validationErr})
		return
	}
	family, err := userModel.CreateFamily(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Family could not be created, admin already has family", "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully created", "family": family})
}

//user add a family(user send request to admin)
func (ctrl UserController) AddFamilyUser(c *gin.Context) {
	var form forms.RequestAddFamilyForm
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "parameter error"})
		return
	}
	if form.Admin_id == 0 || form.User_id == 0 {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "id can not be 0"})
		return
	}
	err := userModel.AddFamilyUser(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Unable to send or too many requests sent", "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sent successfully, waiting for admin response"})
}

// Admin Receive new response GET
func (ctrl UserController) ReceiveNewAdmin(c *gin.Context) {
	//get user id and check the new resoponse list
	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		return
	}

	results, err := userModel.AllAdminNewResponse(getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Could not get new requests"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

// admin receive one response pram:admin_id, person_id, isConfirm
func (ctrl UserController) ConfirmNewFamilyAdmin(c *gin.Context) {
	var form forms.ConfirmRequestAddFamilyForm
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "parameter error"})
		return
	}
	//change reponse db status and update user's famliy_id
	err := userModel.ConfirmUserToFamily(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Could not add new family member into database"})
		return
	}
	if form.Confirm == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "New member rejected"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "New member added successfully"})
}

//Get user list which belong to this family
func (ctrl UserController) GetFamilyMembersByFamilyID(c *gin.Context) {
	id := c.Param("familyID")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		return
	}
	results, err := userModel.FamilyListByFamilyID(getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Could not get family members"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})

}
