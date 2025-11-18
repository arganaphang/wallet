package service

import (
	"context"

	ulid "github.com/oklog/ulid/v2"

	"github.com/arganaphang/wallet/backend/internal/entity"
	"github.com/arganaphang/wallet/backend/internal/repository"
)

type ICategoryService interface {
	Add(ctx context.Context, name string) error
	DeleteByID(ctx context.Context, id ulid.ULID) error
	GetAll(ctx context.Context, query string) ([]entity.Category, error)
}

type categoryService struct {
	repositories repository.Repository
}

func NewCategoryService(repositories repository.Repository) ICategoryService {
	return &categoryService{repositories: repositories}
}

// Add implements ICategoryService.
func (c *categoryService) Add(ctx context.Context, name string) error {
	return c.repositories.Category.Add(ctx, name)
}

// DeleteByID implements ICategoryService.
func (c *categoryService) DeleteByID(ctx context.Context, id ulid.ULID) error {
	return c.repositories.Category.DeleteByID(ctx, id)
}

// GetAll implements ICategoryService.
func (c *categoryService) GetAll(ctx context.Context, query string) ([]entity.Category, error) {
	return c.repositories.Category.GetAll(ctx, query)
}
