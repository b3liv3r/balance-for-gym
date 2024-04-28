package service

import (
	"context"
	"github.com/b3liv3r/balance-for-gym/modules/balance/models"
)

type Walleter interface {
	Create(ctx context.Context, userID int) (string, error)
	GetByID(ctx context.Context, userID int) (models.Wallet, error)
	Update(ctx context.Context, userID int, amount float64, Type, description string) error
	ListTransactions(ctx context.Context, userID int) ([]models.Transaction, error)
}
