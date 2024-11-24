package entities

import (
	"time"

	"github.com/joaofilippe/todoGo/pkg/enum"
)

// Todo represents a task with a title, description, and status.
type Todo struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deadline    time.Time
	Status      enum.Status
}
