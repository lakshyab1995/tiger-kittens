package graph

import (
	"context"
	"encoding/base64"
	"strconv"

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

func EncodeCursor(id int) string {
	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(id)))
}
func DecodeCursor(encoded string) (int, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(decoded))
}
