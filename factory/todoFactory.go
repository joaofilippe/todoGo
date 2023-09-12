package factory

import todoService "todoGo/application/services/Todo"

type TodoFactory struct {
	TodoService    *todoService.TodoService
}

func (f *TodoFactory) Create() *todoService.TodoService {
	return f.TodoService
}