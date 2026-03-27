package repository

import (
	"fmt"
	"merch/internal/model"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllUsers() ([]model.User, error) {
	users := make([]model.User, 0)
	err := r.db.Select(&users, "Select * from users")
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to find users: %w", err)
	}
	return users, nil
}

func (r *Repository) GetUser(username string) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * from users WHERE username=$1", username)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to find user %s: %w", username, err)
	}
	return &user, nil
}

func (r *Repository) CreateUser(user *model.User) (int, error) {
	var id int
	err := r.db.QueryRowx(`INSERT INTO users (username, password_hash) 
		VALUES ($1, $2)`, user.Username, user.PasswordHash).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgres: failed to insert user %s: %w", user.Username, err)
	}
	return id, nil
}
