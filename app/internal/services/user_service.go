package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/models"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/repositories"
)

var (
	ErrUsernameExists = errors.New("username already exists")
	ErrEmailExists    = errors.New("email already exists")
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepository.GetAll()
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.UserRepository.GetByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	if s.UserRepository.UsernameExists(user.Username) {
		return ErrUsernameExists
	}

	if s.UserRepository.EmailExists(user.Email) {
		return ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.UserRepository.Create(user)
}

func (s *UserService) UpdateUser(id string, user *models.User) error {
	return s.UserRepository.Update(id, user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.UserRepository.Delete(id)
}
