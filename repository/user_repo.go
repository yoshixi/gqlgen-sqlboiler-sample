package repository

import (
	"context"
  "database/sql"

	"github.com/yoshixj/gqlgen-sqlboiler-sample/domain"
	"github.com/yoshixj/gqlgen-sqlboiler-sample/infra/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// NewUserRepository ...
func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
    db: db,
  }
}

type userRepository struct {
  db *sql.DB
}

// List ...
func (p *userRepository) List(ctx context.Context) ([]*domain.User, error) {
	users, err := models.Users(qm.Limit(10)).All(ctx, p.db)
	if err != nil {
		return []*domain.User{}, err
	}
  return []*domain.User{ { User: users[0] } }, nil
}


// GetByID ...
func (p *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	u, err := models.FindUserG(ctx, id)
	if err != nil {
		return &domain.User{}, err
	}
  return &domain.User{User: u}, nil
}


// Create ...
func (p *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
  err := user.Insert(ctx, p.db, boil.Infer())

  if err!= nil {
    return &domain.User{}, err
  }

  return user, nil
}


