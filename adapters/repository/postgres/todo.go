package postgres

import (
	application "todoGo/application/domain/models"
)

type TodoRepository struct {
	Connection *Connection
}

func (r *TodoRepository) Create(t application.Todo) (int64, error) {
	// var todo dto.TodoDTO

	// todo = todo.ToTodoRepositoryModel(t)



	return 0, nil
}