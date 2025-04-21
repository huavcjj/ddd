package entity

import "errors"

type User struct {
	name string
}

func NewUser(name string) (User, error) {
	if len(name) < 3 {
		return User{}, errors.New("user name is too short")
	}
	return User{name: name}, nil
}

func (u User) Name() string {
	return u.name
}

// 後からユーザ名の変更を行うことができない
