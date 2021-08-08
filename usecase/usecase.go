package usecase

import (
	"github.com/yoshixj/gqlgen-sqlboiler-sample/domain"
	"github.com/yoshixj/gqlgen-sqlboiler-sample/repository"
)

// NewUsecase ....
func NewUsecase(repo repository.Repository) Usecase {
	return &usecaseImpl{
		repo,
	}
}

// Usecase ...
type Usecase interface {
	User() domain.UserUsecase
}

type usecaseImpl struct {
	repo repository.Repository
}

func (impl usecaseImpl) User() domain.UserUsecase {
	return NewUserUsecase(impl.repo)
}

