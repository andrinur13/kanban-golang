package main

import (
	"fmt"
	"kanban-golang/conf"
	"kanban-golang/controller"
	"kanban-golang/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"kanban-golang/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db := conf.InitDB()

	// repository
	userRepository := repository.NewUserRepository(db)

	// service
	userService := service.NewUserService(userRepository)

	// controller
	userController := controller.NewUserController(userService)

	router := gin.Default()

	// routing
	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.Login)

	router.Run()

	fmt.Println("hello world!")
}
