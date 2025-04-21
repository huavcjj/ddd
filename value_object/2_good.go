package value_object

import "errors"

type UserName struct {
	name string
}

func NewUserName(name string) (UserName, error) {
	if len(name) < 3 {
		return UserName{}, errors.New("user name is too short")
	}
	return UserName{name: name}, nil
}
