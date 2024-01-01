package models

import (
	"time"

	common "github.com/joaofilippe/todoGo/common/enum"
)

type Todo struct {
	Id          int64
	UserId      int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deadline    time.Time
	Status      common.Status
}
