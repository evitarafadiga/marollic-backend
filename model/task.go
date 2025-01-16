package model

type Task struct {
	ID        int    `json:"id_task"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"deleted"`
}
