package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (u UserController) Retrieve(c *gin.Context) {
	var users []User
	var user User
	user.Name = "Bake some cake"
	user.Email = `thiduzz14@gmail.com`

	users = append(users, user)
	c.JSON(http.StatusOK, gin.H{"message": "User founded!", "users": users})
	c.Abort()
	return
}