package services

import (
	"api/internal/models"
	"api/internal/repositories"
	"errors"
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	Create(user UserDTO) (*models.User, error)
	Find() ([]*models.User, error)
	FindByID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(UserDTO UserDTO) (*models.User, error) {
	existingUser, _ := s.repo.FindByEmail(UserDTO.Email)
	if existingUser != nil {
		return nil, errors.New("email já está sendo utilizado")
	}

	user := &models.User{
		ID:        uuid.New().String(),
		Name:      UserDTO.Name,
		Email:     UserDTO.Email,
		Password:  UserDTO.Password,
		CreatedAt: time.Now(),
	}

	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Find() ([]*models.User, error) {
	return s.repo.Find()
}

func (s *userService) FindByID(id string) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) FindByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) Update(user *models.User) error {
	return s.repo.Update(user)
}

func (s *userService) Delete(id string) error {
	return s.repo.Delete(id)
}
