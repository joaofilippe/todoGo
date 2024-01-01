package todo

import todoModels "github.com/joaofilippe/todoGo/application/models/todo"

type TodoService struct {
	TodoService    ITodoService
	TodoRepository ITodoRepository
}

func (s *TodoService) Create(todo *todoModels.Todo) (int, error) {
	todoId, err := s.TodoRepository.Create(todo)
	if err != nil {
		return 0, err
	}

	return todoId, err
}
