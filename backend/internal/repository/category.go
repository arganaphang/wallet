package repository

import (
	"context"
	"fmt"

	"github.com/arganaphang/wallet/backend/internal/entity"
	goqu "github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type ICategoryRepository interface {
	Add(ctx context.Context, name string) error
	GetAll(ctx context.Context, query string) ([]entity.Category, error)
}

type categoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) ICategoryRepository {
	return &categoryRepository{DB: db}
}

// Add implements ICategoryRepository.
func (c *categoryRepository) Add(ctx context.Context, name string) error {
	sql, _, err := goqu.
		Insert(entity.TABLE_CATEGORIES).
		Cols(
			"name",
		).
		Vals(goqu.Vals{
			name,
		}).
		ToSQL()
	if err != nil {
		return err
	}

	_, err = c.DB.Exec(sql)
	return err
}

// GetAll implements ICategoryRepository.
func (c *categoryRepository) GetAll(ctx context.Context, q string) ([]entity.Category, error) {
	query := goqu.From(entity.TABLE_CATEGORIES)

	if q != "" {
		query = query.Where(goqu.C("name").Like(fmt.Sprintf("%%%s%%", q)))
	}

	sql, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	var results []entity.Category
	err = c.DB.Select(&results, sql)

	return results, err
}
