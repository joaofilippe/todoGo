package postgres

import (
	todoModels "github.com/joaofilippe/todoGo/internal/domain/entities/todo"
)

type TodoRepository struct {
	Connection *Connection
}

func (r *TodoRepository) Create(t todoModels.Todo) (int64, error) {
	// var todo dto.TodoDTO

	// todo = todo.ToTodoRepositoryModel(t)



	return 0, nil
}