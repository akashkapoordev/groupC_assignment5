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

// Invitation struct represents an invitation code //B
type Invitation struct {
	ID       int       json:"id"
	Code     string    json:"code"
	Used     bool      json:"used"
	IssuedAt time.Time json:"issued_at"
}



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