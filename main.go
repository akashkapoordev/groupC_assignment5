package main

// Define database connection details
const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "pass123456"
	dbname   = "users"
)

// Define database tables
const (
	usersTable       = "users"
	invitationsTable = "invitations"
	adminsTable      = "admins"
)

// SetupDatabase creates a connection to the PostgreSQL database
func SetupDatabase() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}