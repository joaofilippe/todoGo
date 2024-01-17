package queries

const (
	// SelectUserQuery is the query to select a user by id
	SelectUserQuery = "SELECT * FROM users WHERE id = $1"

	// SelectAllUsersQuery is the query to select all users
	SelectAllUsersQuery = "SELECT * FROM users"

	// InsertNewUserQuery is the query to insert a user
	InsertNewUserQuery = `
		INSERT INTO public.user(
			id, 
			first_name, 
			last_name, 
			username, 
			email, 
			password, 
			birthdate, 
			active
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)
	`
)