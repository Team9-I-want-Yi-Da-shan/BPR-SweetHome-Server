package controllers

import (
	"gin-Home-server/forms"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ArticleController ...
type DeviceController struct{}
 

//Create ...
func (ctrl DeviceController) Create(c *gin.Context) {
	userID := getUserID(c)

	var form forms.CreateArticleForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := articleForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	id, err := articleModel.Create(userID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article created", "id": id})
}

//All ...
func (ctrl DeviceController) All(c *gin.Context) {
	userID := getUserID(c)

	results, err := articleModel.All(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Could not get articles"})
		// return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}
