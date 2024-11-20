package models

import (
	"time"

	"github.com/joaofilippe/todoGo/pkg/enum"
)

type Todo struct {
	Id          int64
	UserId      int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deadline    time.Time
	Status      enum.Status
}
