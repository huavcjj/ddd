package dto

import "ddd/repository/sample/domain"

type UserDTO struct {
	ID   string
	Name string
}

func NewUserDTO(user *domain.User) *UserDTO {
	return &UserDTO{
		ID:   user.ID(),
		Name: user.Name(),
	}
}
