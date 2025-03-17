package services

import (
	"fmt"
	"context"
	"log/slog"

	"github.com/gintokos/coinder/backend/internal/constants"
	"github.com/gintokos/coinder/backend/internal/models"
	"github.com/gintokos/coinder/backend/pkg/gerror"
	"github.com/gintokos/coinder/backend/pkg/sl"
	"github.com/gintokos/coinder/backend/pkg/telegram"
)

type UserStorage interface {
	UpdateUser(user models.User) error
}

type UserService struct {
	storage UserStorage
	bot     *telegram.Bot
}

func NewUserService(bot *telegram.Bot, storage UserStorage) *UserService {
	return &UserService{
		storage: storage,
		bot:     bot,
	}
}

func (s *UserService) Update(user models.User) error {
	err := s.storage.UpdateUser(user)
	if err != nil {
		slog.Error("error on updating user", sl.Err(err))
		return constants.ErrServer
	}
	return nil
}

func (s *UserService) CreateInvoice(amount int) (string, error) {
	invoicestr, err := s.bot.CreateInvoiceLinkStars(context.TODO(), amount)
	if err != nil {
		slog.Error("error on creating invoice", sl.Err(err))
		return "", gerror.New(fmt.Errorf("error on creating invoice: %v", err), constants.ErrServer, 500)
	}
	return invoicestr, nil
}
