package services

import (
	"errors"
	"testing"

	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/mocks"
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/models"
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/utils"

	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {

	hash, _ :=
		utils.HashPassword("123456")

	mockRepo :=
		&mocks.MockUserRepository{
			User: &models.User{
				ID:       1,
				Username: "admin",
				Password: hash,
			},
		}

	authService :=
		NewAuthService(mockRepo)

	token, err :=
		authService.Login(
			"admin",
			"123456",
		)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestLoginWrongPassword(
	t *testing.T,
) {

	hash, _ :=
		utils.HashPassword("123456")

	mockRepo :=
		&mocks.MockUserRepository{
			User: &models.User{
				ID:       1,
				Username: "admin",
				Password: hash,
			},
		}

	authService :=
		NewAuthService(mockRepo)

	token, err :=
		authService.Login(
			"admin",
			"wrongpassword",
		)

	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestLoginUserNotFound(t *testing.T) {

	mockRepo := &mocks.MockUserRepository{
		User: nil,
		Err:  errors.New("record not found"),
	}

	service := NewAuthService(mockRepo)

	token, err := service.Login(
		"admin",
		"123456",
	)

	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestRegisterDuplicateUsername(
	t *testing.T,
) {

	mockRepo := &mocks.MockUserRepository{
		User: &models.User{
			ID:       1,
			Username: "admin",
		},
	}

	service := NewAuthService(mockRepo)

	err := service.Register(
		"admin",
		"123456",
	)

	assert.Error(t, err)

	assert.Equal(
		t,
		"Username already exists",
		err.Error(),
	)
}

func TestRegisterSuccess(
	t *testing.T,
) {

	mockRepo :=
		&mocks.MockUserRepository{}

	authService :=
		NewAuthService(mockRepo)

	err :=
		authService.Register(
			"admin",
			"123456",
		)

	assert.NoError(t, err)
}
