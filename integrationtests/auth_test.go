package integrationtests

import (
	"bytes"
	"encoding/json"
	"go-login-register/configs"
	"go-login-register/controllers"
	"go-login-register/requests"
	"go-login-register/responses"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	configs.ConnectDatabase()
	os.Exit(m.Run())
}

func setupTestDB() *gorm.DB {
	tx := configs.DB.Begin()
	return tx
}

func teardownTestDB(tx *gorm.DB) {
	tx.Rollback()
}

func TestSignUpExpectSuccessful(t *testing.T) {
	router := http.NewServeMux()
	router.HandleFunc("/signup", controllers.SignUp)

	requestBody := requests.SignUpRequest{
		Name:            "John Doe",
		Email:           "john.doe@example.com",
		Password:        "password123",
		PasswordConfirm: "password123",
	}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var response responses.BaseValueResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	assert.Equal(t, "User Registered Successfully", response.Message)
}

func TestSignUpWithMismatchedPasswordsExpectError(t *testing.T) {
	router := http.NewServeMux()
	router.HandleFunc("/signup", controllers.SignUp)

	requestBody := requests.SignUpRequest{
		Name:            "John Doe",
		Email:           "john.doe@example.com",
		Password:        "password123",
		PasswordConfirm: "differentpassword",
	}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	var response responses.BaseValueResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	assert.Equal(t, "Password and Password Confirm Not Match", response.Message)
}

func TestLoginExpectSuccesfull(t *testing.T) {
	router := http.NewServeMux()
	router.HandleFunc("/login", controllers.Login)

	requestBody := requests.LoginRequest{
		Email:    "john.doe@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response responses.BaseValueResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	assert.Equal(t, "Successfully Login", response.Message)
}

func TestSignUpWithMissingFields(t *testing.T) {
	router := http.NewServeMux()
	router.HandleFunc("/signup", controllers.SignUp)

	requestBody := requests.SignUpRequest{}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	var response responses.BaseValueResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	assert.Contains(t, response.Message, "Validation error")
}

func TestLoginWithIncorrectCredentials(t *testing.T) {
	router := http.NewServeMux()
	router.HandleFunc("/login", controllers.Login)

	requestBody := requests.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "wrongpassword",
	}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	var response responses.BaseValueResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	assert.Equal(t, "Email or Password incorrect", response.Message)
}

func TestLoginWithMissingFields(t *testing.T) {
	router := http.NewServeMux()
	router.HandleFunc("/login", controllers.Login)

	requestBody := requests.LoginRequest{}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	var response responses.BaseValueResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	assert.Contains(t, response.Message, "Validation error")
}
