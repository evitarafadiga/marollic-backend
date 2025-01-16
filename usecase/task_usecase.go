package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type TaskUseCase struct {
	repository repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) TaskUseCase {
	return TaskUseCase{
		repository: repo,
	}
}

func (tu *TaskUseCase) GetTasks() ([]model.Task, error) {

	return tu.repository.GetTasks()
}