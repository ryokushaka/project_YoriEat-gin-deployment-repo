package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) domain.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (cr *categoryRepository) Create(ctx context.Context, category *domain.Category) error {
	query := `INSERT INTO categories (name, bg_color, txt_color, image) VALUES ($1, $2, $3, $4)`
	_, err := cr.db.ExecContext(ctx, query, category.Name, category.BgColor, category.TxtColor, category.Image)
	return err
}

func (cr *categoryRepository) Fetch(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category
	query := `SELECT id, name, bg_color, txt_color, image FROM categories`
	err := cr.db.SelectContext(ctx, &categories, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Category{}, nil
		}
		return nil, err
	}
	return categories, nil
}

func (cr *categoryRepository) GetByID(ctx context.Context, id string) (domain.Category, error) {
	var category domain.Category
	query := `SELECT id, name, bg_color, txt_color, image FROM categories WHERE id = $1`
	err := cr.db.GetContext(ctx, &category, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return category, nil
		}
		return category, err
	}
	return category, nil
}
