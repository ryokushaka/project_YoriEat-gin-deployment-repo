package domain

import (
	"context"
)

const (
	CollectionUser = "users"
)

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

type UserUsecase interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
	Update(c context.Context, user *User) error
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
	Update(c context.Context, user *User) error
}
