package repository

import (
  "database/sql"
  "github.com/yoshixj/gqlgen-sqlboiler-sample/domain"
  "github.com/yoshixj/gqlgen-sqlboiler-sample/infra/conn"
)

// NewRepository ...
func NewRepository() (Repository, func()) {
	db, closeDB := conn.SetupDB()
	return repositoryImpl{
    db: db,
  }, func() {
		closeDB()
	}
}

// Repository ...
type Repository interface {
	User() domain.UserRepository
}

type repositoryImpl struct {
  db *sql.DB
}

func (impl repositoryImpl) User() domain.UserRepository {
  return NewUserRepository(impl.db)
}

