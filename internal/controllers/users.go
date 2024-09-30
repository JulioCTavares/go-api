package controllers

import (
	"api/internal/services"
	"encoding/json"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {}
func GetUser(w http.ResponseWriter, r *http.Request)  {}
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body services.UserDTO

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	user, err := uc.userService.Create(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}
