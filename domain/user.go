package domain

import (
  "context"
	"github.com/yoshixj/gqlgen-sqlboiler-sample/infra/models"
)

type User struct {
  *models.User
}

type UserRepository interface {
  List(ctx context.Context) ([]*User, error)
  GetByID(ctx context.Context, id string) (*User, error)
}

type UserUsecase interface {
  List(ctx context.Context) ([]*User, error)
  GetByID(ctx context.Context, id string) (*User, error)
}
