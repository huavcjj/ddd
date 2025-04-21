package value_object

import (
	"errors"
)

type User struct {
	Id       string
	userName UserName
}

type UserName struct {
	name string
}

func NewUserName(name string) (UserName, error) {
	if len(name) < 3 {
		return UserName{}, errors.New("user name is too short")
	}
	return UserName{name: name}, nil
}

func NewUser(id string, userName UserName) (User, error) {
	return User{
		Id:       id,
		userName: userName,
	}, nil
}

func CreateUser(name string) error {
	userName, err := NewUserName(name)
	if err != nil {
		return err
	}
	id := "generated-id"
	_, err = NewUser(id, userName)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(id, name string) error {
	userName, err := NewUserName(name)
	if err != nil {
		return err
	}
	_, err = NewUser(id, userName)
	if err != nil {
		return err
	}
	return nil
}
