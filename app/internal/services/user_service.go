package services

import (
	"errors"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/models"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/repositories"
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrUsernameExists     = errors.New("username already exists")
	ErrEmailExists        = errors.New("email already exists")
)

type UserService struct {
	UserRepository *repositories.UserRepository
	jwtSecret      string
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
		jwtSecret:      os.Getenv("JWT_SECRET"), // Carrega a chave secreta de uma variável de ambiente
	}
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

func (s *UserService) Authenticate(identifier, password string) (string, error) {
	user, err := s.UserRepository.FindByUsernameOrEmail(identifier)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
