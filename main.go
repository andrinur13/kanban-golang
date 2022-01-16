package main

import (
	"fmt"
	"kanban-golang/conf"
	"kanban-golang/controller"
	"kanban-golang/middleware"
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
	taskRepository := repository.NewTaskRepository(db)

	// service
	userService := service.NewUserService(userRepository)
	taskService := service.NewTaskService(taskRepository)

	// controller
	userController := controller.NewUserController(userService)
	taskController := controller.NewTaskController(taskService)

	router := gin.Default()

	// routing

	// user
	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.Login)
	router.PUT("/update-account", middleware.AuthMiddleware(), userController.UpdateUser)

	// task
	router.POST("/tasks", middleware.AuthMiddleware(), taskController.CreateNewTask)
	router.GET("/tasks", middleware.AuthMiddleware(), taskController.GetAllTask)
	router.PUT("/tasks/:id", middleware.AuthMiddleware(), taskController.UpdateTask)
	router.PATCH("/tasks/update-status/:id", middleware.AuthMiddleware(), taskController.UpdateStatusTask)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware(), taskController.DeleteTask)

	// taskSwithced := true
	// a, b := taskRepository.SwitchStatus(3, taskSwithced)
	// fmt.Println(a)
	// fmt.Println(b)

	router.Run()

}
