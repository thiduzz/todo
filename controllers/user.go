package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thiduzz/todo/models"
	"net/http"
)

type UserController struct{}

func (u UserController) Index(c *gin.Context) {
	var users []models.User
	var user models.User
	user.Name = "Bake some cake"
	user.Email = `thiduzz14333@gmail.com`
	users = append(users, user)
	c.JSON(http.StatusOK, gin.H{"message": "User founded!", "users": users})
	return
}