package middleware

import (
	"context"
	"fmt"
	"go/adv-example/configs"
	"go/adv-example/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "email"
)

func IsAuthenticated(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("IsAuthenticated")
		token := r.Header.Get("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)

		req := r.WithContext(ctx)
		fmt.Println(isValid, data)
		next.ServeHTTP(w, req)
	})
}
