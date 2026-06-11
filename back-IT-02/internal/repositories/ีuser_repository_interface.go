package repositories

import "github.com/Rafiana219/interview-question-02/back-IT-02/internal/models"

type IUserRepository interface {
	FindByUsername(username string) (*models.User, error)
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
}
