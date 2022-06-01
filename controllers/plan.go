package controllers

import (
	 
	"net/http"
	"strconv"

	"gin-Home-server/forms"
	"gin-Home-server/models"

	"github.com/gin-gonic/gin"
)

type PlanController struct{}

var planModel = new(models.PlanModel)
var planForm = new(forms.PlanForm)

//Api add new personal plan
func (ctrl PlanController) AddPersonalPlan(c *gin.Context) {
	var personalPlanForm forms.PersonalPlanForm

	if validationErr := c.ShouldBindJSON(&personalPlanForm); validationErr != nil {
		// message := planForm
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"validation Error": validationErr, "current form": personalPlanForm})
		return
	}
	user, err := planModel.CreatePersonalPlan(personalPlanForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error(), "form": personalPlanForm})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully added", "plan_id": user})
}
func (ctrl PlanController) AddFamilyPlan(c *gin.Context) {
	var personalPlanForm forms.FamilyPlanForm

	if validationErr := c.ShouldBindJSON(&personalPlanForm); validationErr != nil {
		// message := planForm
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"validation Error": validationErr, "current form": personalPlanForm})
		return
	}
	user, err := planModel.CreateFamilylPlan(personalPlanForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error(), "form": personalPlanForm})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully added", "plan_id": user})
}

//Api get user's all plan
func (ctrl PlanController) GetPersonalPlanListByPersonID(c *gin.Context) {
	//personal id
	var idform forms.Id
	if err := c.ShouldBindJSON(&idform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error user id": err.Error()})
		return
	}
	results, err := planModel.AllPersonPlan(idform.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error(), "form": idform})
		return
	}
	c.JSON(http.StatusOK, gin.H{"results": results})
}

//Api get family's all plan
func (ctrl PlanController) GetFamilyPlanListByFamilyID(c *gin.Context) {
	//personal id
	var idform forms.Family_Id
	if err := c.ShouldBindJSON(&idform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error user id": err.Error()})
		return
	}
	results, err := planModel.AllFamilyPlan(idform.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error(), "form": idform})
		return
	}
	c.JSON(http.StatusOK, gin.H{"results": results})

}

//API get one plan by plan id
func (ctrl PlanController) GetPersonalPlanByPlanID(c *gin.Context) {
	//plan id
	var idform forms.Plan_Id
	if err := c.ShouldBindJSON(&idform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error plan id": err.Error()})
		return
	}
	result, err := planModel.GetPlanByID(idform.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error(), "form": idform})
		return
	}
	c.JSON(http.StatusOK, gin.H{"plan": result})
}

//API get one plan by plan id
func (ctrl PlanController) GetFamilyPlanByPlanID(c *gin.Context) {
	//plan id
	var idform forms.Plan_Id
	if err := c.ShouldBindJSON(&idform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error plan id": err.Error()})
		return
	}
	result, err := planModel.GetFamilyPlanByID(idform.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error(), "form": idform})
		return
	}
	c.JSON(http.StatusOK, gin.H{"plan": result})
}

//API DELETE personal plan
func (ctrl PlanController) RemovePersonPlanByPlanID(c *gin.Context) {
	var idform forms.Plan_Id
	if err := c.ShouldBindJSON(&idform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error plan id": err.Error()})
		return
	}
	err := planModel.DeletePersonalPlan("plan_person", idform.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "could not find plan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plan deleted"})
}

//API DELETE personal plan
func (ctrl PlanController) RemoveFamilyPlanByPlanID(c *gin.Context) {
	var idform forms.Family_Id
	if err := c.ShouldBindJSON(&idform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error plan id": err.Error()})
		return
	}
	err := planModel.DeleteFamilyPlan(idform.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "could not find plan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plan deleted"})
}

//API Update personal plan
func (ctrl PlanController) UpdatePersonPlanByPlanID(c *gin.Context) {

	// userID := getUserID(c)

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid parameter", "plan_id": id})
		return
	}
	var personalPlanForm forms.PersonalPlanForm

	if validationErr := c.ShouldBindJSON(&personalPlanForm); validationErr != nil {
		// message := planForm
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"validation Error": validationErr, "current form": personalPlanForm})
		return
	}
	erro := planModel.UpdatePersonPlan(getID, personalPlanForm)
	if erro != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "plan id not exit", "error": erro})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully update"})
}

//API Update personal plan
func (ctrl PlanController) UpdateFamilyPlanByPlanID(c *gin.Context) {

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid parameter", "plan_id": id})
		return
	}
	var familyPlanForm forms.FamilyPlanForm

	if validationErr := c.ShouldBindJSON(&familyPlanForm); validationErr != nil {
		// message := planForm
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"validation Error": validationErr, "current form": familyPlanForm})
		return
	}
	erro := planModel.UpdateFamilyPlan(getID, familyPlanForm)
	if erro != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "plan id not exit", "error": erro})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully update"})
}
