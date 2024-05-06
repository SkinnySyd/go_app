package usercontroller

import (
	"ginhello/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var user []models.User

	models.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{"users": user})

}
