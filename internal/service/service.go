package service

import (
	"database/sql"
	"errors"
	"fmt"
	"merch/internal/model"
	"merch/internal/repository"
	"merch/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	rep *repository.Repository
}

func NewService(rep *repository.Repository) *Service {
	return &Service{rep: rep}
}

func (s *Service) GetAllUsers() ([]model.User, error) {
	users, err := s.rep.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("service: failed to get all users: %w", err)
	}
	return users, nil
}

func (s *Service) Auth(uname, passwd string) (*string, error) {
	user, err := s.rep.GetUser(uname)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			hash, err := bcrypt.GenerateFromPassword([]byte(passwd), 10)
			if err != nil {
				return nil, fmt.Errorf("service: couldn't calculate hash %w", err)
			}
			user := &model.User{
				Username:     uname,
				PasswordHash: string(hash),
			}
			id, errCreate := s.rep.CreateUser(user)
			if errCreate != nil {
				return nil, fmt.Errorf("service: couldn't create user %w", errCreate)
			}
			jwt, errJwt := jwt.GenerateToken(id)
			if errJwt != nil {
				return nil, fmt.Errorf("service: couldn't generate jwt: %w", errJwt)
			}
			return &jwt, nil
		} else {
			return nil, fmt.Errorf("service: database error %s: %w", uname, err)
		}
	}
	err2 := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(passwd))
	if err2 != nil {
		return nil, fmt.Errorf("service: wrong password")
	}
	jwt, err3 := jwt.GenerateToken(user.ID)
	if err3 != nil {
		return nil, fmt.Errorf("service: couldn't generate jwt: %w", err3)
	}
	return &jwt, nil
}
