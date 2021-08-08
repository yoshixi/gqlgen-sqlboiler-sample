package usecase

import (
	"context"

	"github.com/yoshixj/gqlgen-sqlboiler-sample/domain"
	"github.com/yoshixj/gqlgen-sqlboiler-sample/repository"
)

// UserUsecase ...
type UserUsecase struct {
	repo repository.Repository
}

// NewUserUsecase ..
func NewUserUsecase(repo repository.Repository) domain.UserUsecase {
	return &UserUsecase{repo}
}

// List ...
func (u UserUsecase) List(ctx context.Context) ([]*domain.User, error) {
  return u.repo.User().List(ctx)
}

// GetByID  ...
func (u UserUsecase) GetByID(ctx context.Context, id string) (*domain.User, error) {
  return u.repo.User().GetByID(ctx, id)
}

