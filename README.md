# Go Login Register Service - Technical Test 

## Overview

This is a RESTful API application written in Go, which allows users to sign up, log in, and retrieve their profile information. This service uses a MySQL database for storage and GORM for ORM. JWT tokens are used for authentication.

## Project Structure

```
go-login-register
 ┣ configs
 ┃ ┗ database.go  #initialize connection to MySQL Database
 ┣ controllers
 ┃ ┣ authController.go  #contains logic for sign up and login
 ┃ ┗ userController.go  #contains logic for current logged in user data retrieval
 ┣ dto
 ┃ ┗ profileDto.go  #define profile dto
 ┣ integrationtests
 ┃ ┣ auth_test.go  #login and signup integration tests
 ┃ ┗ user_test.go  #get current logged in user data integration tests
 ┣ middlewares
 ┃ ┗ auth.go  #middleware to check for jwt token
 ┣ models
 ┃ ┗ user.go  #define user model
 ┣ requests
 ┃ ┣ loginRequest.go  #define login request body
 ┃ ┗ signUpRequest.go  #define signup request body
 ┣ responses
 ┃ ┗ response.go  #define the structure of responses
 ┣ routes
 ┃ ┣ auth.go  #define route and api prefix for auth controller
 ┃ ┗ user.go  #define route and api prefix for user controller
 ┣ utils
 ┃ ┣ password.go  #contains logic to hashpassword and verify
 ┃ ┗ token.go  #contains logic for generate jwt token and verify
 ┣ .env  #env var file(not included in version control)
 ┣ .env.example #env var file example
 ┣ .gitignore
 ┣ README.md
 ┣ go.mod
 ┣ go.sum
 ┗ main.go  #entry point of application
```


## Installation

To set up the project on your local machine, follow these steps:

1. **Clone the Repository**

```bash
git clone https://github.com/andrewtanjaya/go-login-register.git
cd go-login-register
```

2. **Set Up Environment Variables**

Copy the .env.example file to .env and configure the necessary environment variables. For example:

```bash
cp .env.example .env
```

Edit .env with your environment-specific settings. It might include variables like:

```
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=pass
DB_PORT=3306
DB_NAME=testdb

JWT_SECRET_KEY=secret
```

3. **Install Dependencies**

Use `go mod` to download the required dependencies:

```bash
go mod tidy
```

4. **Run the Application**

Use the following command to run the service:

```bash
go run main.go
```

The application will start on port `8080` by default. You can access it at `http://localhost:8080`


## API Endpoints

### SignUp
* Endpoint : `POST /api/auth/register`
* Description : Register a new user
* Request Body :
```json
{
    "name": "andrewtan",
    "email": "andrewtanjaya13@gmail.com",
    "password": "admin",
    "password_confirm": "admin"
}
```
* Response :
```json
{
    "code": 201,
    "message": "User Registered Successfully"
}
```

### Login
* Endpoint : `POST /api/auth/login`
* Description : Logs in a user and returns a JWT token.
* Request Body :
```json
{
    "email": "andrewtanjaya13@gmail.com",
    "password": "admin"
}
```
* Response :
```json
{
    "code": 200,
    "message": "Successfully Login",
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAsIm5hbWUiOiJhbmRyZXd0YW4iLCJlbWFpbCI6ImFuZHJld3RhbmpheWExM0BnbWFpbC5jb20iLCJleHAiOjE3MjM3NTkxNTAsIm5iZiI6MTcyMzc1MTk1MCwiaWF0IjoxNzIzNzUxOTUwfQ.CaMtu1FtgCh3t-5wXwZj6RYDPIeJClmnD4Pf1QMX4vY"
}
```


### Me
* Endpoint : `GET /api/users/me`
* Description : Retrieves the profile of the currently logged-in user.
* Headers :
```json
Authorization: Bearer your_jwt_token
```
* Response :
```json
{
    "code": 200,
    "message": "Current Login User Data",
    "data": {
        "id": 10,
        "name": "andrewtan",
        "email": "andrewtanjaya13@gmail.com"
    }
}
```

## Running Tests

To run the integrations tests, use: 

```bash
go test -v ./integrationtests
```

Ensure that the test environment is correctly configured with a test database.


