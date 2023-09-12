package models

import (
	"time"
)

type Todo struct {
	Id          int64     `gorm:"id"`
	UserId      int64     `gorm:"user_id"`
	Title       string    `gorm:"title"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
	Deadline    time.Time `gorm:"deadline"`
	Status      string    `gorm:"status"`
}
