package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thiduzz/todo/models"
	"github.com/thiduzz/todo/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/users/", UserController.Retrieve)
	r.Run()
}