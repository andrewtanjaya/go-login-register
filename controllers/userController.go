package controllers

import (
	"go-login-register/dto"
	"go-login-register/responses"
	"go-login-register/utils"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value("current_user").(*utils.JWTClaims)
	response := &dto.ProfileDto{
		ID:    currentUser.ID,
		Name:  currentUser.Name,
		Email: currentUser.Email,
	}
	responses.Response(w, 200, "Current Login User Data", response)
}
