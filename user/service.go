package user

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	GetUserByID(ID uuid.UUID) (User, error)
	LoginUser(input LoginInput) (User, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Username = input.UserName
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.Email = input.Email
	user.Name = input.Name

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) GetUserByID(ID uuid.UUID) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == uuid.Nil {
		return user, errors.New("User Not Found")
	}
	return user, nil
}
func (s *service) GetAllUser() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}
func (s *service) LoginUser(input LoginInput) (User, error) {
	username := input.UserName
	password := input.Password

	user, err := s.repository.FindByUserName(username)
	if err != nil {
		return user, err
	}
	if user.ID == uuid.Nil {
		return user, errors.New("User Not Found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, err
}
