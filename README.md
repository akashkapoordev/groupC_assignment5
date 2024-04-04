# User Management API

This project is a simple RESTful API for user management, including registration, login, invitation code generation, and admin registration.

## Requirements

- Go programming language (version 1.16 or later)
- PostgreSQL database
- Dependencies listed in `go.mod` file
- cURL or any REST API client for testing

## Setup

1. **Clone the repository:**

   ```go
   git clone https://github.com/akashkapoordev/groupC_assignment4.git
   ```

2. **Install dependencies:**
    ```go
    cd groupC_assignment4
    go mod download
    ```

3. **Set up the PostgreSQL database:**

- Create a new PostgreSQL database named users.
- For database, please check "Database script" folder, inside this folder you will find postgres script.
- Please run that script, makesure your Docker is running.
- Update the database connection details in the main.go file if necessary.

4. **Build and run the application:**
    ```go
    go run main.go
    ```
The server will start running at http://localhost:8081

**Endpoints**
- User Registration:
    http://localhost:8081/register
    ```go
    curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"username": "example_user", "password": "example_password", "invitation_code": "example_invitation_code"}' \
  
    ```
- User Login:
    http://localhost:8081/login
    ```go
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"username": "example_user", "password": "example_password"}' \
    
    ```
    - Generate Invitation Code:
      http://localhost:8081/generate-invitation
       ```go
       curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"username": "admin_username", "password": "admin_password"}' \
       ```
    - Register Admin:
    http://localhost:8081/register-admin
    ```go
    curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"username": "new_admin", "password": "admin_password"}' \

    ```

    **To check web pages, please visit below endpoints on any browser.**
    For login/register page:
    
    ```bash
    http://localhost:8081/
    ```

    For Invite page, here you have to provide admin credentials to get invitation code.
    ```bash
    http://localhost:8081/invite
    ```



    -
