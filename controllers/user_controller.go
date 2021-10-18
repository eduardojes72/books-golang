package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/database"
	"github.com/hyperyuri/webapi-with-go/models"
	"github.com/hyperyuri/webapi-with-go/services"
)

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind json: " + err.Error(),
		})
		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.Status(204)

}
