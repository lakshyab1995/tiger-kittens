package graph

import (
	"context"

	"github.com/lakshyab1995/tiger-kittens/graph/model"
	"github.com/lakshyab1995/tiger-kittens/jwt"
)

func RefreshToken(ctx context.Context, token string) (*model.TokenMeta, error) {
	username, err := jwt.ParseToken(token)
	if err != nil {
		return nil, err
	}
	jwtToken, err := jwt.GenerateToken(username)
	if err != nil {
		return nil, err
	}
	return &model.TokenMeta{
		Token:  jwtToken.Token,
		Expiry: jwtToken.Expiry,
	}, nil
}
