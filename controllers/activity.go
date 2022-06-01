package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-Home-server/forms"
	"gin-Home-server/models"
)

type ActivityController struct{}

var activityModel = new(models.ActivityModel)
var activityForm = new(forms.ActivityForm)

func (ctrl ActivityController) CreateFamilyActivity(c *gin.Context) {
	var form forms.AllActivityForm
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": validationErr, "form": form})
		return
	}

	id, err := activityModel.Create(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "activity could not be created", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "activity created", "activity_id": id})
}

//update  activity
func (ctrl ActivityController) UpdateActivity(c *gin.Context) {
	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid parameter", "id": id})
		return
	}
	var form forms.AllActivityForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		// message := planForm
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"validation Error": validationErr, "current form": form})
		return
	}
	erro := activityModel.Update(getID, form)
	if erro != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "plan id not exit", "error": erro})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully update"})
}

//TODO: Get activities by familyID and date
func (ctrl ActivityController) GetActivityByFamilyDate(c *gin.Context) {

	var form forms.ActivityDataFamilyForm
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": validationErr, "form": form})
		return
	}
	results, err := activityModel.FamilyIDDate(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Could not get activities by date and id", "err": err.Error()})
	}

	c.JSON(http.StatusOK, results)

}

//Get activities by familyID and date
func (ctrl ActivityController) GetActivityByPersonDate(c *gin.Context) {

	var form forms.ActivityDataPersonForm
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": validationErr, "form": form})
		return
	}
	results, err := activityModel.PersonIDDate(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Could not get activities by date and id", "err": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})

}

//Get one activity by id
func (ctrl ActivityController) GetActivityByID(c *gin.Context) {
	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid parameter", "acivity_id": id})
		return
	}
	result, erro := activityModel.GetActivityByID(getID)
	if erro != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": erro.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"activity": result})
}

// finish activity ,
func (ctrl ActivityController) SetFinish(c *gin.Context) {
	var idform forms.Activity_Id
	if err := c.ShouldBindJSON(&idform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error plan id": err.Error()})
		return
	}
	err := activityModel.SetFinish(idform.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "could not find activity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity finish"})
}

//delete
func (ctrl ActivityController) RemoveActivityByPlanID(c *gin.Context) {
	var idform forms.Activity_Id
	if err := c.ShouldBindJSON(&idform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error plan id": err.Error()})
		return
	}
	err := activityModel.DeleteActivityByID(idform.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "could not find activity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted"})
}
