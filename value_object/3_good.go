package value_object

type UserId struct {
	id string
}

type UserName struct {
	name string
}

type User struct {
	Id   UserId
	Name UserName
}

func CreateUser(name UserName) (User, error) {
	var user User
	user.Id = name //コンパイルエラーになる
	return user, nil
}
