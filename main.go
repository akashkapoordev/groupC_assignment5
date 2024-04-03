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

// RegisterHandler handles user registration 
func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse request body to extract user details
		var requestBody struct {
			Username       string json:"username"
			Password       string json:"password"
			InvitationCode string json:"invitation_code"
		}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		// Extract user details from the request body
		username := requestBody.Username
		password := requestBody.Password
		invitationCode := requestBody.InvitationCode

		// Validate invitation code
		if invitationCode == "" {
			http.Error(w, "Invitation code is required", http.StatusBadRequest)
			return
		}
		valid, err := validateInvitationCode(db, invitationCode)
		if err != nil {
			http.Error(w, "Failed to validate invitation code", http.StatusInternalServerError)
			return
		}
		if !valid {
			http.Error(w, "Invalid invitation code", http.StatusUnauthorized)
			return
		}

		// Check if the username already exists
		exists, err := isUsernameExists(db, username)
		if err != nil {
			http.Error(w, "Failed to check username existence", http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}

		// Hash user password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		// Save user details to the database
		_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, string(hashedPassword))
		if err != nil {
			http.Error(w, "Failed to save user", http.StatusInternalServerError)
			return
		}

		// Mark invitation code as used
		err = markInvitationCodeAsUsed(db, invitationCode)
		if err != nil {
			log.Println("Failed to mark invitation code as used:", err)
			// This is a non-critical error, so continue with the registration
		}

		fmt.Fprintf(w, "User registered successfully")
	}
}