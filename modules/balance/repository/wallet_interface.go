package repository

import (
	"context"
	"github.com/b3liv3r/balance-for-gym/modules/balance/models"
)

type WalleterRepository interface {
	Create(ctx context.Context, userID int) error
	GetByID(ctx context.Context, userID int) (models.Wallet, error)
	Update(ctx context.Context, userID int, amount float64) error
	AddTransaction(ctx context.Context, transaction models.Transaction) error
	ListTransactionsByUser(ctx context.Context, userID int) ([]models.Transaction, error)
}
