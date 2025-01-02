package entity

import (
	"time"

	ulid "github.com/oklog/ulid/v2"
)

type TransactionType string

const (
	INCOME  TransactionType = "income"
	OUTCOME TransactionType = "outcome"
)

const TABLE_TRANSACTIONS = "transactions"

type Transaction struct {
	ID        ulid.ULID       `json:"id"         db:"id"`
	Name      string          `json:"name"       db:"name"`
	Amount    uint64          `json:"amount"     db:"amount"`
	Category  string          `json:"category"   db:"category"`
	Type      TransactionType `json:"type"       db:"type"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
}
