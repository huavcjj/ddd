package value_object

import "errors"

// CheckUserName 値を利用する前にルールに照らし合わせる必要がある
func CheckUserName(userName string) error {
	if len(userName) >= 3 {
		return nil
	}
	return errors.New("user name is too short")
}
