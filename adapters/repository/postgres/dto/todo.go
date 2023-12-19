package dto

import (
	"time"
	application "github.com/joaofilippe/todoGo/application/domain/models"
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

func (d *TodoDTO) ToTodoRepositoryModel(t application.Todo) {
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