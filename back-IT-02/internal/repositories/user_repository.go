package repositories

import (
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/config"
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/models"
)

type IUserRepository interface {
	FindByUsername(username string) (*models.User, error)
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByUsername(
	username string,
) (*models.User, error) {

	var user models.User

	err := config.DB.
		Where("username = ?", username).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(
	user *models.User,
) error {

	return config.DB.Create(user).Error
}

func (r *UserRepository) FindByID(
	id uint,
) (*models.User, error) {

	var user models.User

	err := config.DB.
		First(&user, id).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
