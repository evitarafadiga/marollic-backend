package controller

import (
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