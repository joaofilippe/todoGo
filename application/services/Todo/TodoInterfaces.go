package Todo

import "todoGo/application/domain/models"

type ITodoService interface {
	Create(todo *models.Todo) (int, error)
}

type ITodoRepository interface {
	Create(todo *models.Todo) (int, error)
}
