package integrationtests

import (
	"encoding/json"
	"go-login-register/controllers"
	"go-login-register/dto"
	"go-login-register/middlewares"
	"go-login-register/models"
	"go-login-register/responses"
	"go-login-register/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMe(t *testing.T) {
	testUser := models.User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	token, err := utils.CreateToken(&testUser)
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	router := http.NewServeMux()
	router.Handle("/me", middlewares.Auth(http.HandlerFunc(controllers.Me)))

	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response responses.BaseValueResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	assert.Equal(t, "Current Login User Data", response.Message)

	profileData, ok := response.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected response.Data to be map[string]interface{}, but got %T", response.Data)
	}

	profileJSON, err := json.Marshal(profileData)
	if err != nil {
		t.Fatalf("Failed to marshal profile data: %v", err)
	}

	var profile dto.ProfileDto
	err = json.Unmarshal(profileJSON, &profile)
	if err != nil {
		t.Fatalf("Failed to unmarshal profile data: %v", err)
	}

	assert.Equal(t, testUser.ID, profile.ID)
	assert.Equal(t, testUser.Name, profile.Name)
	assert.Equal(t, testUser.Email, profile.Email)
}
