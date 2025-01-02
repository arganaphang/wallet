package repository

import (
	"context"

	goqu "github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	ulid "github.com/oklog/ulid/v2"

	"github.com/arganaphang/wallet/backend/internal/entity"
)

type ITransactionRepository interface {
	Add(ctx context.Context, data entity.Transaction) error
	GetAll(ctx context.Context) ([]entity.Transaction, error)
	GetByID(ctx context.Context, id ulid.ULID) (*entity.Transaction, error)
	UpdateByID(ctx context.Context, id ulid.ULID, data entity.Transaction) error
	DeleteByID(ctx context.Context, id ulid.ULID) error
}

type transactionRepo struct {
	DB *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) ITransactionRepository {
	return &transactionRepo{DB: db}
}

// Add implements ITransactionRepository.
func (t *transactionRepo) Add(ctx context.Context, data entity.Transaction) error {
	sql, _, err := goqu.Insert(entity.TABLE_TRANSACTIONS).
		Cols(
			"id",
			"name",
			"amount",
			"category",
			"type",
		).Vals(goqu.Vals{
		ulid.Make().String(),
		data.Name,
		data.Amount,
		data.Category,
		data.Type,
	}).ToSQL()
	if err != nil {
		return err
	}

	_, err = t.DB.Exec(sql)
	return err
}

// DeleteByID implements ITransactionRepository.
func (t *transactionRepo) DeleteByID(ctx context.Context, id ulid.ULID) error {
	sql, _, err := goqu.
		Delete(entity.TABLE_TRANSACTIONS).
		Where(goqu.C("id").Eq(id.String())).
		ToSQL()
	if err != nil {
		return err
	}

	_, err = t.DB.Exec(sql)

	return err
}

// GetTransactionByID implements ITransactionRepository.
func (t *transactionRepo) GetByID(ctx context.Context, id ulid.ULID) (*entity.Transaction, error) {
	sql, _, err := goqu.
		From("transactions").
		Where(goqu.C("id").Eq(id.String())).
		Limit(1).
		ToSQL()
	if err != nil {
		return nil, err
	}

	var result entity.Transaction
	err = t.DB.Get(&result, sql)

	return &result, err
}

// GetTransactions implements ITransactionRepository.
func (t *transactionRepo) GetAll(ctx context.Context) ([]entity.Transaction, error) {
	sql, _, err := goqu.From(entity.TABLE_TRANSACTIONS).ToSQL()
	if err != nil {
		return nil, err
	}

	var results []entity.Transaction
	err = t.DB.Select(&results, sql)

	return results, err
}

// UpdateByID implements ITransactionRepository.
func (t *transactionRepo) UpdateByID(ctx context.Context, id ulid.ULID, data entity.Transaction) error {
	sql, _, err := goqu.
		Update(entity.TABLE_TRANSACTIONS).
		Set(goqu.Record{
			"name":     data.Name,
			"amount":   data.Amount,
			"category": data.Category,
			"type":     data.Type,
		}).
		Where(goqu.C("id").Eq(id.String())).
		ToSQL()
	if err != nil {
		return err
	}

	_, err = t.DB.Exec(sql)

	return err
}
