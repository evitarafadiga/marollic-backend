package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if (err != nil) {
		panic(err)
	}

	//Camada de repository
	TaskRepository := repository.NewTaskRepository(dbConnection)
	//Camada usecase
	TaskUseCase := usecase.NewTaskUseCase(TaskRepository)
	//Camada de Controllers
	TaskController := controller.NewTaskController(TaskUseCase)
	
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/tasks", TaskController.GetTasks)

	server.Run(":8000")
}