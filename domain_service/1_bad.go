package domain_service

import (
	"errors"
	"github.com/google/uuid"
)

type UserID struct {
	id string
}

type User struct {
	userID   UserID
	userName string
}

func NewUserID() UserID {
	id := uuid.NewString()
	return UserID{id: id}
}

func NewUser(userID UserID, userName string) (*User, error) {
	if len(userID.String()) == 0 {
		return nil, errors.New("user ID is empty")
	}
	if len(userName) < 3 {
		return nil, errors.New("user name is too short")
	}
	return &User{
		userID:   userID,
		userName: userName,
	}, nil
}

func (u UserID) String() string {
	return u.id
}

// Exists 生成したオブジェクト自身に問い合わせをすることになる
func (u *User) Exists(other *User) bool {
	return u.userID.String() == other.userID.String()
}
