package mocks

import (
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/models"
)

type MockUserRepository struct {
	User *models.User
	Err  error
}

func (m *MockUserRepository) FindByUsername(
	username string,
) (*models.User, error) {

	return m.User, m.Err
}

func (m *MockUserRepository) Create(
	user *models.User,
) error {

	return m.Err
}

func (m *MockUserRepository) FindByID(
	id uint,
) (*models.User, error) {

	return m.User, m.Err
}
