package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/domain"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (name, email, password, tags, bio, social, image) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := ur.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Tags, user.Bio, user.Social, user.Image)
	return err
}

func (ur *userRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	query := `SELECT id, name, email, tags, bio, social, image FROM users`
	err := ur.db.SelectContext(ctx, &users, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.User{}, nil
		}
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	query := `SELECT id, name, email, password, tags, bio, social, image FROM users WHERE email = $1`
	err := ur.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func (ur *userRepository) GetByID(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	query := `SELECT id, name, email, password, tags, bio, social, image FROM users WHERE id = $1`
	err := ur.db.GetContext(ctx, &user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func (ur *userRepository) Update(ctx context.Context, user *domain.User) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3, tags = $4, bio = $5, social = $6, image = $7 WHERE id = $8`
	_, err := ur.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Tags, user.Bio, user.Social, user.Image, user.ID)
	return err
}
