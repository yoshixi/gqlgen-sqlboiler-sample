package domain

import (
	"context"
	"github.com/yoshixj/gqlgen-sqlboiler-sample/infra/models"
)

// User ...
type User struct {
	*models.User
}

// UserRepository ...
type UserRepository interface {
	List(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
}

// UserUsecase ...
type UserUsecase interface {
	List(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
}
