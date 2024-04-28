package service

import (
	"context"
	"github.com/b3liv3r/balance-for-gym/modules/balance/models"
	"github.com/b3liv3r/balance-for-gym/modules/balance/repository"
	"go.uber.org/zap"
	"time"
)

type WalletService struct {
	repo   repository.WalleterRepository
	logger *zap.Logger
}

func NewWalletService(repo repository.WalleterRepository, logger *zap.Logger) Walleter {
	return &WalletService{
		repo:   repo,
		logger: logger,
	}
}

func (ws *WalletService) Create(ctx context.Context, userID int) (string, error) {
	err := ws.repo.Create(ctx, userID)
	if err != nil {
		return "", err
	}
	return "Wallet created successfully", nil
}

func (ws *WalletService) GetByID(ctx context.Context, userID int) (models.Wallet, error) {
	wallet, err := ws.repo.GetByID(ctx, userID)
	if err != nil {
		return models.Wallet{}, err
	}
	return wallet, nil
}

func (ws *WalletService) Update(ctx context.Context, userID int, amount float64, ttype, description string) error {
	err := ws.repo.Update(ctx, userID, amount)
	if err != nil {
		return err
	}

	// Добавляем транзакцию в историю
	transaction := models.Transaction{
		UserId:      userID,
		Amount:      amount,
		Type:        ttype,
		Description: description,
		Date:        time.Now(),
	}
	err = ws.repo.AddTransaction(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (ws *WalletService) ListTransactions(ctx context.Context, userID int) ([]models.Transaction, error) {
	transactions, err := ws.repo.ListTransactionsByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
