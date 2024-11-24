package dto

import (
	"time"

	"github.com/joaofilippe/todoGo/internal/domain/entities"
)

// TodoDTO represents a Data Transfer Object for Todo.
type TodoDTO struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deadline    time.Time
	Status      string
}

// FromEntity converts an entities.Todo to a TodoDTO.
func (d *TodoDTO) FromEntity(t entities.Todo) {
	d.ID = t.ID
	d.UserID = t.UserID
	d.Title = t.Title
	d.Description = t.Description
	d.CreatedAt = t.CreatedAt
	d.UpdatedAt = t.UpdatedAt
	d.Deadline = t.Deadline
	d.Status = t.Status.ToString()

}
