package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

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

	var task model.Task
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

func (t *taskController) GetTaskById(ctx *gin.Context) {

	id := ctx.Param("taskId")
	if id == "" {
		response := model.Response{
			Message: "Id da tarefa não pode ser nula.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id da tarefa precisa ser um número.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	task, err := t.TaskUseCase.GetTaskById(taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if task == nil {
		response := model.Response{
			Message: "Tarefa não encontrada.",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, task)
}
