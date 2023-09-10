package main

import (
	"todoGo/application/services/Todo"
	"todoGo/application/services/User"
)

type application struct {
	TodoService *todo.TodoService
	UserService *user.UserService
}

func main() {

}
