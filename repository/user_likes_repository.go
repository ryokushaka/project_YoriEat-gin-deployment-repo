package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type userLikesRepository struct {
	db *sqlx.DB
}

func NewUserLikesRepository(db *sqlx.DB) domain.UserLikesRepository {
	return &userLikesRepository{
		db: db,
	}
}

func (ulr *userLikesRepository) AddLike(ctx context.Context, userID, recipeID int) error {
	query := `INSERT INTO user_likes (user_id, recipe_id) VALUES ($1, $2)`
	_, err := ulr.db.ExecContext(ctx, query, userID, recipeID)
	return err
}

func (ulr *userLikesRepository) RemoveLike(ctx context.Context, userID, recipeID int) error {
	query := `DELETE FROM user_likes WHERE user_id = $1 AND recipe_id = $2`
	_, err := ulr.db.ExecContext(ctx, query, userID, recipeID)
	return err
}

func (ulr *userLikesRepository) FetchLikesByUserID(ctx context.Context, userID int) ([]domain.UserLikes, error) {
	var likes []domain.UserLikes
	query := `SELECT user_id, recipe_id FROM user_likes WHERE user_id = $1`
	err := ulr.db.SelectContext(ctx, &likes, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.UserLikes{}, nil
		}
		return nil, err
	}
	return likes, nil
}
