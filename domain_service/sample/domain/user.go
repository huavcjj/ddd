package domain

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrUserNameTooShort = errors.New("user name is too short")
)

type UserID struct {
	value string
}

func NewUserID() UserID {
	return UserID{value: uuid.NewString()}
}

type UserName struct {
	value string
}

func NewUserName(name string) UserName {
	return UserName{value: name}
}

type User struct {
	UserID   UserID
	UserName UserName
}

func NewUser(userName UserName) (*User, error) {
	if len(userName.value) < 3 {
		return nil, ErrUserNameTooShort
	}

	return &User{
		UserID:   NewUserID(),
		UserName: userName,
	}, nil
}

func (u *UserID) String() string {
	return u.value
}

func (u *UserName) String() string {
	return u.value
}
