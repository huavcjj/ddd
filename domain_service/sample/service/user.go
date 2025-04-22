package service

import (
	"ddd/domain_service/sample/domain"
	"errors"
	"sync"
)

var userStore sync.Map

type userService struct {
}

type UserService interface {
	CreateUser(name string) (*domain.User, error)
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) CreateUser(name string) (*domain.User, error) {

	if _, exists := userStore.Load(name); exists {
		return nil, errors.New("user already exists")
	}

	userName := domain.NewUserName(name)
	user, err := domain.NewUser(userName)
	if err != nil {
		return nil, err
	}

	userStore.Store(name, user)
	return user, nil
}
