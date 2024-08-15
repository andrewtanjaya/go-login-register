package controllers

import (
	"encoding/json"
	"go-login-register/configs"
	"go-login-register/models"
	"go-login-register/requests"
	"go-login-register/responses"
	"go-login-register/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func SignUp(w http.ResponseWriter, r *http.Request) {
	var request requests.SignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responses.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	errValidate := validate.Struct(request)
	if errValidate != nil {
		responses.Response(w, 400, "Validation error: "+errValidate.Error(), nil)
		return
	}

	if request.Password != request.PasswordConfirm {
		responses.Response(w, 400, "Password and Password Confirm Not Match", nil)
		return
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		responses.Response(w, 500, err.Error(), nil)
		return
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		responses.Response(w, 500, err.Error(), nil)
		return
	}

	responses.Response(w, 201, "User Registered Successfully", nil)

}

func Login(w http.ResponseWriter, r *http.Request) {
	var request requests.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responses.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	errValidate := validate.Struct(request)
	if errValidate != nil {
		responses.Response(w, 400, "Validation error: "+errValidate.Error(), nil)
		return
	}

	var user models.User
	if err := configs.DB.First(&user, "email = ?", request.Email).Error; err != nil {
		responses.Response(w, 400, "Email or Password incorrect", nil)
		return
	}

	if err := utils.VerifyPassword(user.Password, request.Password); err != nil {
		responses.Response(w, 400, "Email or Password incorrect", nil)
		return
	}

	token, err := utils.CreateToken(&user)
	if err != nil {
		responses.Response(w, 500, err.Error(), nil)
		return
	}

	responses.Response(w, 200, "Successfully Login", token)
}
