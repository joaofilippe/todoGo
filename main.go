package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func main() {
	dbConn, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=12345678 dbname=todo sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	users := []User{}
	err = dbConn.Select(&users, "SELECT * FROM teste")
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		println(user.ID)
		println(user.Name)
		println(user.Email)
	}
}
