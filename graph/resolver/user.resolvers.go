package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
  "fmt"

	"github.com/yoshixj/gqlgen-sqlboiler-sample/domain"
	"github.com/yoshixj/gqlgen-sqlboiler-sample/graph/model"
	"github.com/yoshixj/gqlgen-sqlboiler-sample/infra/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*domain.User, error) {
	user := createUserInputToDomain(input)
	err := r.uc.User().Create(ctx, user)
	return user, err
}

func (r *queryResolver) Users(ctx context.Context) (*model.UserConnection, error) {
	list, err := r.uc.User().List(ctx)
	return userListToConnection(list), err
}

func (r *queryResolver) User(ctx context.Context, id string) (*domain.User, error) {
  return r.uc.User().GetByID(ctx, id)
}

func createUserInputToDomain(input model.CreateUserInput) *domain.User {
	u := &models.User{FirstName: input.FirstName, LastName: input.LastName, Email: input.Email}
	return &domain.User{
		User: u,
	}
}
func userListToConnection(users []*domain.User) *model.UserConnection {
  fmt.Printf("%v", users)

	edges := make([]*model.UserEdge, len(users))
  fmt.Printf("\n fmt %+v", users[0].FirstName)
  for i, v := range users {
    edges[i] = &model.UserEdge{ Node: v, Cursor: ""}
  }
	return &model.UserConnection{
		PageInfo: &model.PageInfo{},
		Edges:    edges,
	}
}
