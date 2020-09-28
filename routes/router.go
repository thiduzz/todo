package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thiduzz/todo/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		userController := new(controllers.UserController)
		v1.GET("/users", userController.Index)

	}

	return router

}