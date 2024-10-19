package domain

import (
	"context"
)

const (
	// CollectionUser is the name of the user collection in the database.
	CollectionUser = "users"
)

// User represents a user entity.
type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Tags     []string `json:"tags"`
	Bio      string   `json:"bio"`
	Social   []string `json:"social"`
	Image    string   `json:"image"`
	Password string   `json:"password"`
}

// UserUsecase represents the user use case interface.
type UserUsecase interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
	Update(c context.Context, user *User) error
}

// UserRepository represents the user repository interface.
type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
	Update(c context.Context, user *User) error
}
