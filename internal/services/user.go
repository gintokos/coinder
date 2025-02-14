package services

import (
	"log/slog"

	"github.com/gintokos/coinder/internal/constants"
	"github.com/gintokos/coinder/internal/models"
	"github.com/gintokos/coinder/pkg/sl"
)

type UserStorage interface {
	UpdateUser(user models.User) error
}

type UserService struct {
	storage UserStorage
}

func NewUserService(storage UserStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (s UserService) Update(user models.User) error {
	err := s.storage.UpdateUser(user)
	if err != nil {
		slog.Error("error on updating user", sl.Err(err))
		return constants.ErrServer
	}
	return nil
}
