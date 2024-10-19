package domain

import (
	"context"
)

// Category represents a category entity.
type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	BgColor  string `json:"bg_color"`
	TxtColor string `json:"txt_color"`
	Image    string `json:"image"`
}

// CategoryUsecase represents the category use case interface.
type CategoryUsecase interface {
	Create(c context.Context, category *Category) error
	Fetch(c context.Context) ([]Category, error)
	GetByID(c context.Context, id string) (Category, error)
}

// CategoryRepository represents the category repository interface.
type CategoryRepository interface {
	Create(c context.Context, category *Category) error
	Fetch(c context.Context) ([]Category, error)
	GetByID(c context.Context, id string) (Category, error)
}
