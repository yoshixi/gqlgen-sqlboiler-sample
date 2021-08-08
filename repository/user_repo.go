package repository

import (
	"context"
  "database/sql"
  "fmt"

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
  fmt.Printf("List repository")
	users, err := models.Users(qm.Limit(10)).All(ctx, p.db)
	if err != nil {
		return []*domain.User{}, err
	}
  list := make([]*domain.User, len(users))

  for i, u := range users {
    list[i] = &domain.User{ User: u }
  }

  return list, nil
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
func (p *userRepository) Create(ctx context.Context, user *domain.User) error {
  err := user.Insert(ctx, p.db, boil.Infer())

  return err
}


