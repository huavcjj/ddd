package domain

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrUserNameIsTooShort = errors.New("user name is too short")
	ErrUserNameIsTooLong  = errors.New("user name is too long")
	ErrUserIDEmpty        = errors.New("user ID is empty")
)

type UserID struct {
	value string
}

func NewUserID(value string) (UserID, error) {
	if len(value) == 0 {
		return UserID{}, ErrUserIDEmpty
	}
	return UserID{value: value}, nil
}

func (id UserID) String() string {
	return id.value
}

type UserName struct {
	value string
}

func NewUserName(value string) (UserName, error) {
	if len(value) < 3 {
		return UserName{}, ErrUserNameIsTooShort
	}
	if len(value) > 20 {
		return UserName{}, ErrUserNameIsTooLong
	}
	return UserName{value: value}, nil
}

func (n UserName) String() string {
	return n.value
}

type User struct {
	id   UserID
	name UserName
}

func NewUser(name UserName) (*User, error) {
	id, err := NewUserID(uuid.NewString())
	if err != nil {
		return nil, err
	}
	return &User{
		id:   id,
		name: name,
	}, nil
}

func NewUserFromStrings(idStr, nameStr string) (*User, error) {
	id, err := NewUserID(idStr)
	if err != nil {
		return nil, err
	}
	name, err := NewUserName(nameStr)
	if err != nil {
		return nil, err
	}
	return &User{
		id:   id,
		name: name,
	}, nil
}

func (u *User) ID() string {
	return u.id.String()
}

func (u *User) Name() string {
	return u.name.String()
}
