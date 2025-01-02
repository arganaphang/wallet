package service

import (
	"context"

	ulid "github.com/oklog/ulid/v2"

	"github.com/arganaphang/wallet/backend/internal/entity"
	"github.com/arganaphang/wallet/backend/internal/repository"
)

type ITransactionService interface {
	Add(ctx context.Context, data entity.Transaction) error
	GetAll(ctx context.Context) ([]entity.Transaction, error)
	GetByID(ctx context.Context, id ulid.ULID) (*entity.Transaction, error)
	UpdateByID(ctx context.Context, id ulid.ULID, data entity.Transaction) error
	DeleteByID(ctx context.Context, id ulid.ULID) error
}

type transactionService struct {
	repositories repository.Repository
}

func NewTransactionService(repositories repository.Repository) ITransactionService {
	return &transactionService{repositories: repositories}
}

// Add implements ITransactionService.
func (t *transactionService) Add(ctx context.Context, data entity.Transaction) error {
	return t.repositories.Transaction.Add(ctx, data)
}

// DeleteByID implements ITransactionService.
func (t *transactionService) DeleteByID(ctx context.Context, id ulid.ULID) error {
	return t.repositories.Transaction.DeleteByID(ctx, id)
}

// GetTransactionByID implements ITransactionService.
func (t *transactionService) GetByID(ctx context.Context, id ulid.ULID) (*entity.Transaction, error) {
	return t.repositories.Transaction.GetByID(ctx, id)
}

// GetTransactions implements ITransactionService.
func (t *transactionService) GetAll(ctx context.Context) ([]entity.Transaction, error) {
	return t.repositories.Transaction.GetAll(ctx)
}

// UpdateByID implements ITransactionService.
func (t *transactionService) UpdateByID(ctx context.Context, id ulid.ULID, data entity.Transaction) error {
	return t.repositories.Transaction.UpdateByID(ctx, id, data)
}
