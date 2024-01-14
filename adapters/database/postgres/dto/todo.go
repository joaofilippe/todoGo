package dto

import (
	"time"

	todoModels "github.com/joaofilippe/todoGo/application/models/todo"
)

type TodoDTO struct {
	Id          int64
	UserId      int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deadline    time.Time
	Status      string
}

func (d *TodoDTO) ToTodoRepositoryModel(t todoModels.Todo) {
	d = &TodoDTO{
		Id:          t.Id,
		UserId:      t.UserId,
		Title:       t.Title,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		Deadline:    t.Deadline,
		Status:      t.Status.ToString(),
	}
}