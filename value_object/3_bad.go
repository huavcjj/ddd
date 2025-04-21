package value_object

type User struct {
	Id   string
	Name string
}

// CreateUser このコードの正当性はコードを見ただけではわからない
func CreateUser(name string) (User, error) {
	var user User
	user.Id = name // IDにユーザ名を使っているが、他の値を使う場合は間違いになる
	return user, nil
}
