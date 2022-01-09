package service

import (
	"errors"
	"kanban-golang/model/entity"
	"kanban-golang/model/input"
	"kanban-golang/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(userInput input.RegisterUserInput) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(input input.RegisterUserInput) (entity.User, error) {
	newUser := entity.User{}
	newUser.Email = input.Email
	newUser.FullName = input.FullName
	newUser.Role = "member"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return entity.User{}, err
	}

	newUser.Password = string(passwordHash)

	checkSameUser, err := s.userRepository.CheckSameEmail(input.Email)

	if err != nil {
		return entity.User{}, err
	}

	if checkSameUser.ID != 0 {
		return entity.User{}, errors.New("Email already registered!")
	}

	createdUser, err := s.userRepository.Save(newUser)

	if err != nil {
		return entity.User{}, err
	}

	return createdUser, nil
}
