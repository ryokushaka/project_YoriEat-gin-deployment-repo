package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type commentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) domain.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (cr *commentRepository) Create(ctx context.Context, comment *domain.Comment) error {
	query := `INSERT INTO comments (text, user_id, recipe_id, comment_id) VALUES ($1, $2, $3, $4)`
	_, err := cr.db.ExecContext(ctx, query, comment.Text, comment.UserID, comment.RecipeID, comment.CommentID)
	return err
}

func (cr *commentRepository) FetchByRecipeID(ctx context.Context, recipeID int) ([]domain.Comment, error) {
	var comments []domain.Comment
	query := `SELECT id, text, user_id, recipe_id, comment_id FROM comments WHERE recipe_id = $1`
	err := cr.db.SelectContext(ctx, &comments, query, recipeID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Comment{}, nil
		}
		return nil, err
	}
	return comments, nil
}

func (cr *commentRepository) GetByID(ctx context.Context, id string) (domain.Comment, error) {
	var comment domain.Comment
	query := `SELECT id, text, user_id, recipe_id, comment_id FROM comments WHERE id = $1`
	err := cr.db.GetContext(ctx, &comment, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return comment, nil
		}
		return comment, err
	}
	return comment, nil
}
