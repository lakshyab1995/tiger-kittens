package auth

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/lakshyab1995/tiger-kittens/db"
	"github.com/lakshyab1995/tiger-kittens/graph"
	"github.com/lakshyab1995/tiger-kittens/jwt"
	"github.com/lakshyab1995/tiger-kittens/utils"
)

func Middleware(rsv *graph.Resolver) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var ctx context.Context
			header := r.Header.Get("Authorization")
			bearerToken := strings.Split(header, " ")

			if header == "" || len(bearerToken) != 2 {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := bearerToken[1]
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
			log.Println("username from token: ", username)

			// create user and check if user exists in db
			user := db.User{Username: username}
			id, err := rsv.UserRepository.GetUsrIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = id
			log.Println("user id: ", user.ID)
			log.Println("username: ", user.Username)
			// put it in context
			ctx = context.WithValue(r.Context(), utils.UserCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
