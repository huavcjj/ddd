package domain_service

import (
	"errors"
	"github.com/google/uuid"
)

type UserID struct {
	id string
}

func NewUserID() UserID {
	return UserID{id: uuid.NewString()}
}

func (id UserID) String() string {
	return id.id
}

type User struct {
	userID UserID
	name   string
}

func NewUser(userID UserID, name string) (*User, error) {
	if len(userID.String()) == 0 {
		return nil, errors.New("user ID is empty")
	}
	if len(name) < 3 {
		return nil, errors.New("user name is too short")
	}
	return &User{
		userID: userID,
		name:   name,
	}, nil
}

// ChangeName エンティティ側に名前変更の責務を持たせる
func (u *User) ChangeName(newName string) error {
	if len(newName) < 3 {
		return errors.New("user name is too short")
	}
	u.name = newName
	return nil
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}
