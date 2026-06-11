package services

import (
	"errors"

	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/models"
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/repositories"
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/utils"
)

type AuthService struct {
	userRepo repositories.IUserRepository
}

func NewAuthService(
	userRepo repositories.IUserRepository,
) *AuthService {

	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(
	username string,
	password string,
) error {

	if len(username) < 3 {
		return errors.New(
			"username must be at least 3 characters",
		)
	}

	if len(password) < 6 {
		return errors.New(
			"password must be at least 6 characters",
		)
	}

	existingUser, _ := s.userRepo.FindByUsername(username)

	if existingUser != nil {
		return errors.New("Username already exists")
	}

	hash, err := utils.HashPassword(password)

	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: hash,
	}

	return s.userRepo.Create(&user)
}

func (s *AuthService) Login(
	username string,
	password string,
) (string, error) {

	user, err :=
		s.userRepo.FindByUsername(username)

	if err != nil {
		return "", errors.New("not find this username")
	}

	if !utils.CheckPassword(
		user.Password,
		password,
	) {
		return "", errors.New("invalid password")
	}

	token, err :=
		utils.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) GetProfile(
	userID uint,
) (*models.User, error) {

	return s.userRepo.FindByID(userID)
}
