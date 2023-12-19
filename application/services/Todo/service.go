package todo

import "github.com/joaofilippe/todoGo/application/domain/models"

type TodoService struct {
	TodoService    ITodoService
	TodoRepository ITodoRepository
}

func (s *TodoService) Create(todo *models.Todo) (int, error) {
	todoId, err := s.TodoRepository.Create(todo)
	if err != nil {
		return 0, err
	}

	return todoId, err
}
