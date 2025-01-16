package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type TaskRepository struct {
	connection *sql.DB
}

func NewTaskRepository(connection *sql.DB) TaskRepository {
	return TaskRepository{
		connection: connection,
	}
}

func (tr *TaskRepository) GetTasks() ([]model.Task, error) {
	query := "SELECT id, task_name, is_deleted FROM tasks"
	rows, err := tr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Task{}, err

	}

	var taskList []model.Task
	var taskObj model.Task

	for rows.Next() {
		err = rows.Scan(
			&taskObj.ID,
			&taskObj.Name,
			&taskObj.IsDeleted)
		if err != nil {
			return []model.Task{}, err
		}

		taskList = append(taskList, taskObj)
	}

	rows.Close()

	return taskList, nil
}

func (tr *TaskRepository) CreateTask(task model.Task) (int, error) {

	var id int
	query, err := tr.connection.Prepare("INSERT INTO tasks" +
		"(task_name, is_deleted)"+
	" VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(task.Name, task.IsDeleted).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}
