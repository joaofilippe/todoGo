package postgres

import (
	"fmt"

	application "todoGo/application/domain/models"
)

type TodoRepository struct {
	Connection *Connection
}

func (r *TodoRepository) Create(t application.Todo) (int64, error) {
	fmt.Println("Create")

	getLog()
	getOutro()
	getMaisUm()

	return 0, nil
}

func getLog(){
	fmt.Println("colocando um log aqui")
}

func getOutro(){
	fmt.Println("colocando um outro log aqui")
}

func getMaisUm(){
	fmt.Println("colocando um mais um log aqui")
}