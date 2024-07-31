package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type scriptRepository struct {
	db *sqlx.DB
}

func NewScriptRepository(db *sqlx.DB) domain.ScriptRepository {
	return &scriptRepository{
		db: db,
	}
}

func (sr *scriptRepository) Create(ctx context.Context, script *domain.Script) error {
	query := `INSERT INTO scripts (name, user_id) VALUES ($1, $2)`
	_, err := sr.db.ExecContext(ctx, query, script.Name, script.UserID)
	return err
}

func (sr *scriptRepository) Fetch(ctx context.Context) ([]domain.Script, error) {
	var scripts []domain.Script
	query := `SELECT id, name, user_id FROM scripts`
	err := sr.db.SelectContext(ctx, &scripts, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Script{}, nil
		}
		return nil, err
	}
	return scripts, nil
}

func (sr *scriptRepository) GetByID(ctx context.Context, id string) (domain.Script, error) {
	var script domain.Script
	query := `SELECT id, name, user_id FROM scripts WHERE id = $1`
	err := sr.db.GetContext(ctx, &script, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return script, nil
		}
		return script, err
	}
	return script, nil
}

func (sr *scriptRepository) AddRecipeToScript(ctx context.Context, scriptID, recipeID int) error {
	query := `INSERT INTO script_recipes (script_id, recipe_id) VALUES ($1, $2)`
	_, err := sr.db.ExecContext(ctx, query, scriptID, recipeID)
	return err
}

func (sr *scriptRepository) RemoveRecipeFromScript(ctx context.Context, scriptID, recipeID int) error {
	query := `DELETE FROM script_recipes WHERE script_id = $1 AND recipe_id = $2`
	_, err := sr.db.ExecContext(ctx, query, scriptID, recipeID)
	return err
}
