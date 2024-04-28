package repository

import (
	"context"
	"fmt"
	"github.com/b3liv3r/balance-for-gym/modules/balance/models"
	"github.com/jmoiron/sqlx"
)

type WalletRepositoryDB struct {
	db *sqlx.DB
}

func NewWalletRepositoryDB(db *sqlx.DB) WalleterRepository {
	return &WalletRepositoryDB{
		db: db,
	}
}

func (wr *WalletRepositoryDB) Create(ctx context.Context, userID int) error {
	_, err := wr.db.ExecContext(ctx, "INSERT INTO users_wallets (id, balance) VALUES ($1, $2)", userID, 0.0)
	if err != nil {
		return err
	}
	return nil
}

func (wr *WalletRepositoryDB) GetByID(ctx context.Context, userID int) (models.Wallet, error) {
	var wallet models.Wallet
	err := wr.db.GetContext(ctx, &wallet, "SELECT * FROM users_wallets WHERE id = $1", userID)
	if err != nil {
		return models.Wallet{}, err
	}
	return wallet, nil
}
func (wr *WalletRepositoryDB) Update(ctx context.Context, userID int, amount float64) error {
	// Проверяем, есть ли пользователь с таким ID в таблице users_wallets
	var currentBalance float64
	err := wr.db.GetContext(ctx, &currentBalance, "SELECT balance FROM users_wallets WHERE id = $1", userID)
	if err != nil {
		return err
	}

	// Проверяем, достаточно ли средств для списания
	if currentBalance+amount < 0 {
		return fmt.Errorf("insufficient funds")
	}

	// Обновляем баланс
	_, err = wr.db.ExecContext(ctx, "UPDATE users_wallets SET balance = balance + $1 WHERE id = $2", amount, userID)
	if err != nil {
		return err
	}
	return nil
}

func (wr *WalletRepositoryDB) AddTransaction(ctx context.Context, transaction models.Transaction) error {
	_, err := wr.db.ExecContext(ctx, "INSERT INTO transactions_history (user_id, amount, transaction_type, description, transaction_date) VALUES ($1, $2, $3, $4, $5)",
		transaction.UserId, transaction.Amount, transaction.Type, transaction.Description, transaction.Date)
	if err != nil {
		return err
	}
	return nil
}

func (wr *WalletRepositoryDB) ListTransactionsByUser(ctx context.Context, userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := wr.db.SelectContext(ctx, &transactions, "SELECT * FROM transactions_history WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
