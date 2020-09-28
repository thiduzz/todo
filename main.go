package main

import (
	"fmt"
	"github.com/thiduzz/todo/config"
	"github.com/thiduzz/todo/database"
	"github.com/thiduzz/todo/routes"
)

func main() {

	if err := config.Load("config/config.yaml"); err != nil {
		fmt.Println("Failed to load configuration")
		return
	}

	_, err := database.InitDB()
	if err != nil {
		fmt.Println("err open databases")
		return
	}

	router := routes.InitRouter()
	router.Run(config.Get().Addr)
}