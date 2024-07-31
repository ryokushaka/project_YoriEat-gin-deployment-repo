package domain

import (
	"context"
)

type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	BgColor  string `json:"bg_color"`
	TxtColor string `json:"txt_color"`
	Image    string `json:"image"`
}

type CategoryUsecase interface {
	Create(c context.Context, category *Category) error
	Fetch(c context.Context) ([]Category, error)
	GetByID(c context.Context, id string) (Category, error)
}

type CategoryRepository interface {
	Create(c context.Context, category *Category) error
	Fetch(c context.Context) ([]Category, error)
	GetByID(c context.Context, id string) (Category, error)
}
