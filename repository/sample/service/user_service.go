package service

import (
	"ddd/repository/sample/domain"
	"ddd/repository/sample/repository"
	"errors"
)

type userService struct {
	userRepository repository.IUserRepository
}

type IUserService interface {
	CreateUser(name string) (*domain.User, error)
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(name string) (*domain.User, error) {
	existingUser, err := s.userRepository.FindByName(name)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	userName := domain.NewUserName(name)
	user, err := domain.NewUser(userName)
	if err != nil {
		return nil, err
	}

	err = s.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
