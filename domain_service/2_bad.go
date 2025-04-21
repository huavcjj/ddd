package domain_service

import (
	"errors"
	"github.com/google/uuid"
)

type UserID struct {
	id string
}
type User struct {
	userID UserID
	name   string
}

func NewUserID() UserID {
	return UserID{id: uuid.NewString()}
}

func (id UserID) String() string {
	return id.id
}

func NewUser(userID UserID, name string) (*User, error) {
	if len(userID.String()) == 0 {
		return nil, errors.New("user ID is empty")

	}
	if len(name) < 3 {
		return nil, errors.New("user name is too short")
	}
	return &User{
		userID: userID,
		name:   name,
	}, nil
}

type UserService struct{}

type IUserService interface {
	ChangeUserName(userID string, newName string) error
}

func NewUserService() *UserService {
	return &UserService{}
}

// ChangeUserName Domain Serviceに全ての振る舞いを記述するとEntityの役割を果たさなくなる
func (s *UserService) ChangeUserName(userID string, newName string) error {
	// DBに問い合わせる処理を実装する
	// ここでは単純にエラーを返す
	return nil
}
