package entity

import (
	"errors"
	"github.com/google/uuid"
)

type UserID struct {
	value string
}

// NewUserID Value Objectとしての設計
func NewUserID() UserID {
	return UserID{value: uuid.NewString()}
}

func (id UserID) String() string {
	return id.value
}

type User struct {
	userID UserID
	name   string
}

// NewUser Entityとしての設計
func NewUser(name string) (User, error) {
	userID := NewUserID()
	if len(name) < 3 {
		return User{}, errors.New("user name is too short")
	}
	return User{
		userID: userID,
		name:   name,
	}, nil
}

func (u *User) ChangeUserName(name string) error {
	if len(name) < 3 {
		return errors.New("user name is too short")
	}
	u.name = name
	return nil
}

func (u *User) CheckIsSameName(name, otherName string) bool {
	return u.name == otherName
}
