package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type taskController struct {
	TaskUseCase usecase.TaskUseCase
}

func NewTaskController(usecase usecase.TaskUseCase) taskController {
	return taskController{
		TaskUseCase: usecase,
	}
}

func (t *taskController) GetTasks(ctx *gin.Context) {

	tasks, err := t.TaskUseCase.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (t *taskController) CreateTask(ctx *gin.Context) {
	
	var	task model.Task
	err := ctx.BindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedTask, err := t.TaskUseCase.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedTask)
}