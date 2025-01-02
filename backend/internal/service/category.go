package service

import (
	"context"

	"github.com/arganaphang/wallet/backend/internal/entity"
	"github.com/arganaphang/wallet/backend/internal/repository"
)

type ICategoryService interface {
	Add(ctx context.Context, name string) error
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

// GetAll implements ICategoryService.
func (c *categoryService) GetAll(ctx context.Context, query string) ([]entity.Category, error) {
	return c.repositories.Category.GetAll(ctx, query)
}
