package resolver

import "github.com/yoshixj/gqlgen-sqlboiler-sample/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// NewResolver ...
func NewResolver(uc usecase.Usecase) *Resolver {
	return &Resolver{
		uc,
	}
}

// Resolver ...
type Resolver struct {
	uc usecase.Usecase
}
