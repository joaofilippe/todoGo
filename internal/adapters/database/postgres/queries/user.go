package queries

const (
	// SelectUserQueryByID is the query to select a user by id
	SelectUserQueryByID = "SELECT * FROM public.user WHERE id = $1"

	// SelectUserQueryByUsername is the query to select a user by username
	SelectUserQueryByUsername = "SELECT * FROM public.user WHERE username = $1"

	// SelectUserQueryByEmail is the query to select a user by email
	SelectUserQueryByEmail = "SELECT * FROM public.user WHERE email = $1"

	// SelectAllUsersQuery is the query to select all users
	SelectAllUsersQuery = "SELECT * FROM public.user"

	// CheckUserExists
	CheckUserExists = "SELECT EXISTS(SELECT 1 FROM public.user WHERE username = $1 OR email = $2)"

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