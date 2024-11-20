package users_migrations

const (
	// CreateUserTableQuery is the query to create the user table
	CreateUserTableQuery = `
			CREATE TABLE IF NOT EXISTS public.users
		(
		    id         UUID                                NOT NULL
		        CONSTRAINT id_user_pk
		            PRIMARY KEY,
		    first_name VARCHAR                             NOT NULL,
		    last_name  VARCHAR                             NOT NULL,
		    username   VARCHAR                             NOT NULL,
		    email      VARCHAR                             NOT NULL,
		    password   VARCHAR                             NOT NULL,
		    birthdate  DATE                                NOT NULL,
		    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		    updated_at TIMESTAMP,
		    active     BOOL      DEFAULT FALSE             NOT NULL
		);
	`	

	// DropUserTableQuery is the query to drop the user table
	DropUserTableQuery = "DROP TABLE IF EXISTS user"
)