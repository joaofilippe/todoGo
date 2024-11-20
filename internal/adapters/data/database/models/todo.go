package models

import (
	"time"
)

type Todo struct {
	ID          int64     `db:"id"`
	UserID     int64     `db:"user_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Deadline    time.Time `db:"deadline"`
	Status      string    `db:"status"`
}
