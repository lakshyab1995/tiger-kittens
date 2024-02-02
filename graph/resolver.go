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
	UserRepository  db.UserRepository
	TigerRepository db.TigerRepository
	SightRepository db.SightingRepository
}

func NewResolver(gDB *gorm.DB) *Resolver {
	sightRepo := db.NewSightingRepository(gDB)
	return &Resolver{
		UserRepository:  db.NewUserRepository(gDB),
		TigerRepository: db.NewTigerRepository(gDB, sightRepo),
		SightRepository: sightRepo,
	}
}
