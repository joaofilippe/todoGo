package query

const (
	// SelectUserQuery is the query to select a user by id
	SelectUserQuery = "SELECT * FROM users WHERE id = $1"

	// SelectAllUsersQuery is the query to select all users
	SelectAllUsersQuery = "SELECT * FROM users"

	// InsertUserQuery is the query to insert a user
	InsertUserQuery = "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"

	// CreateUserTableQuery is the query to create the user table
	CreateUserTableQuery = `
			CREATE TABLE IF NOT EXISTS public.user
		(
		    id         UUID                                NOT NULL
		        CONSTRAINT id_user_pk
		            PRIMARY KEY,
		    first_name VARCHAR                             NOT NULL,
		    last_name  VARCHAR                             NOT NULL,
		    username   VARCHAR                             NOT NULL,
		    email      VARCHAR                             NOT NULL,
		    password   VARCHAR                             NOT NULL,
		    todo_ids   UUID[],
		    birthdate  DATE                                NOT NULL,
		    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		    updated_at TIMESTAMP,
		    active     BOOL      DEFAULT FALSE             NOT NULL
		);
	`	

	// DropUserTableQuery is the query to drop the user table
	DropUserTableQuery = "DROP TABLE IF EXISTS user"
)