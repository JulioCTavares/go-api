package repositories

import (
	"api/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	Find() ([]*models.User, error)
	FindByID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Find() ([]*models.User, error) {
	var users []*models.User
	return users, r.db.Find(&users).Error
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	return &user, r.db.Where("id = ?", id).First(&user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	return &user, r.db.Where("email = ?", email).First(&user).Error
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id string) error {
	return r.db.Delete(&models.User{}, id).Error
}
