package models

import "time"

type Transaction struct {
	Id          int       `db:"transaction_id"`
	UserId      int       `db:"user_id" `
	Amount      float64   `db:"amount"`
	Type        string    `db:"transaction_type" `
	Description string    `db:"description"`
	Date        time.Time `db:"transaction_date"`
}

type Wallet struct {
	UserID  int     `db:"id" `
	Balance float64 `db:"balance" `
}
