package todo

import todoModels "github.com/joaofilippe/todoGo/application/models/todo"

type ITodoService interface {
	Create(todo *todoModels.Todo) (int, error)
}

type ITodoRepository interface {
	Create(todo *todoModels.Todo) (int, error)
}

