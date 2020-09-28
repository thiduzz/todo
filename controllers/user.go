package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thiduzz/todo/database"
	"github.com/thiduzz/todo/models"
	"net/http"
)

type UserController struct{}

func (u UserController) Index(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"message": "User founded!", "users": users})
	return
}

func (u UserController) Store(c *gin.Context) {
	user := models.User{Name: "Thiago", Email: "thiduzz14@gmail.com", Password: "thithi"}

	_ = database.DB.Create(&user) // pass pointer of data to Create


	c.JSON(http.StatusOK, gin.H{"id": user.ID})
	return
}