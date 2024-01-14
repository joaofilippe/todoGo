package query

const (
	// SelectUserQuery is the query to select a user by id
	SelectUserQuery = "SELECT * FROM users WHERE id = $1"

	// SelectAllUsersQuery is the query to select all users
	SelectAllUsersQuery = "SELECT * FROM users"

	// InsertUserQuery is the query to insert a user
	InsertUserQuery = "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
)