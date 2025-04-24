package service

import (
	"ddd/repository/sample/domain"
	"ddd/repository/sample/dto"
	"ddd/repository/sample/repository"
	"errors"
)

type userService struct {
	userRepository repository.IUserRepository
}

type IUserService interface {
	CreateUser(name string) (*dto.UserDTO, error)
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(name string) (*dto.UserDTO, error) {
	userName, err := domain.NewUserName(name)
	if err != nil {
		return nil, err
	}
	existingUser, err := s.userRepository.FindByName(userName)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}
	user, err := domain.NewUser(userName)
	if err != nil {
		return nil, err
	}

	err = s.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return dto.NewUserDTO(user), nil
}
