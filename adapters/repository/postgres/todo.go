package postgres

import (
	"todoGo/adapters/repository/postgres/dto"
	application "todoGo/application/domain/models"
)

type TodoRepository struct {
	Connection *Connection
}

func (r *TodoRepository) Create(application.Todo t) (int64, error) {
	var todo dto.TodoDTO

	todo = todo.ToTodoRepositoryModel(t)



	return 0, nil
}