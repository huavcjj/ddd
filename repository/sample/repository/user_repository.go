package repository

import (
	"database/sql"
	"ddd/repository/sample/domain"
	"errors"
)

type userRepository struct {
	db *sql.DB
}

type IUserRepository interface {
	FindByName(name string) (*domain.User, error)
	Save(user *domain.User) error
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByName(name string) (*domain.User, error) {
	var id string
	var userName string
	query := `
		SELECT id, name
		FROM users
		WHERE name = ?
	`
	row := r.db.QueryRow(query, name)
	err := row.Scan(&id, &userName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	user, err := domain.NewUserFromStrings(id, userName)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Save(user *domain.User) error {
	existingUser, err := r.FindByName(user.UserName.String())
	if err != nil {
		return err
	}

	if existingUser == nil {
		query := `INSERT INTO users (id, name) VALUES (?, ?)`
		_, err = r.db.Exec(query, user.UserID.String(), user.UserName.String())
		if err != nil {
			return err
		}
	} else {
		query := `UPDATE users SET id = ? WHERE name = ?`
		_, err = r.db.Exec(query, user.UserID.String(), user.UserName.String())
		if err != nil {
			return err
		}
	}

	return nil
}
