package graph

import (
	"github.com/lakshyab1995/tiger-kittens/db"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//go:generate go get -u github.com/99designs/gqlgen@v0.17.43
//go:generate go run github.com/99designs/gqlgen@v0.17.43 generate

type Resolver struct {
	UserRepository db.UserRepository
}

func NewResolver(gDB *gorm.DB) *Resolver {
	return &Resolver{
		UserRepository: db.NewUserRepository(gDB),
	}
}
