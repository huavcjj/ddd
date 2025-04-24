package repository

import (
	"database/sql"
	"ddd/repository/sample/domain"
)

type userRepository struct {
	db *sql.DB
}

func (u userRepository) FindByID(userID domain.UserID) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindByName(userName domain.UserName) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Save(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Delete(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

type IUserRepository interface {
	FindByID(userID domain.UserID) (*domain.User, error)
	FindByName(userName domain.UserName) (*domain.User, error)
	Save(user *domain.User) error
	Delete(user *domain.User) error
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db: db}
}
