package entity

import "errors"

type User struct {
	name string
}

func NewUser(name string) (*User, error) {
	if len(name) < 3 {
		return nil, errors.New("user name is too short")
	}
	return &User{name: name}, nil
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) error {
	if len(name) < 3 {
		return errors.New("user name is too short")
	}
	u.name = name
	return nil
}

// 可能な限り値オブジェクトを利用する
