package value_object

import "errors"

type User struct {
	Id   string
	Name string
}

func CreateUser(name string) (User, error) {
	if len(name) < 3 {
		return User{}, errors.New("user name is too short")
	}
	var user User
	// Do something

	return user, nil
}

func UpdateUser(id, name string) (User, error) {
	if len(name) < 3 {
		return User{}, errors.New("user name is too short")
	}
	var user User
	// Do something

	return user, nil
}

// ユーザ名の最小文字数を変更するときに、ユーザ名の長さ確認を行っているコードを全て変更しないといけない
