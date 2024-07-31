package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type recipeRepository struct {
	db *sqlx.DB
}

func NewRecipeRepository(db *sqlx.DB) domain.RecipeRepository {
	return &recipeRepository{
		db: db,
	}
}

func (rr *recipeRepository) Create(ctx context.Context, recipe *domain.Recipe) error {
	query := `INSERT INTO recipes (name, text, ingredient, time, process, tags, description, category_id, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := rr.db.ExecContext(ctx, query, recipe.Name, recipe.Text, recipe.Ingredient, recipe.Time, recipe.Process, recipe.Tags, recipe.Description, recipe.CategoryID, recipe.UserID)
	return err
}

func (rr *recipeRepository) Fetch(ctx context.Context) ([]domain.Recipe, error) {
	var recipes []domain.Recipe
	query := `SELECT id, name, text, ingredient, time, process, tags, description, category_id, user_id FROM recipes`
	err := rr.db.SelectContext(ctx, &recipes, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Recipe{}, nil
		}
		return nil, err
	}
	return recipes, nil
}

func (rr *recipeRepository) GetByID(ctx context.Context, id string) (domain.Recipe, error) {
	var recipe domain.Recipe
	query := `SELECT id, name, text, ingredient, time, process, tags, description, category_id, user_id FROM recipes WHERE id = $1`
	err := rr.db.GetContext(ctx, &recipe, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return recipe, nil
		}
		return recipe, err
	}
	return recipe, nil
}

func (rr *recipeRepository) Update(ctx context.Context, recipe *domain.Recipe) error {
	query := `UPDATE recipes SET name = $1, text = $2, ingredient = $3, time = $4, process = $5, tags = $6, description = $7, category_id = $8, user_id = $9 WHERE id = $10`
	_, err := rr.db.ExecContext(ctx, query, recipe.Name, recipe.Text, recipe.Ingredient, recipe.Time, recipe.Process, recipe.Tags, recipe.Description, recipe.CategoryID, recipe.UserID, recipe.ID)
	return err
}

func (rr *recipeRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM recipes WHERE id = $1`
	_, err := rr.db.ExecContext(ctx, query, id)
	return err
}
