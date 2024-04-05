package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yunbaifan/go-mall/admin/internal/config"
	"github.com/yunbaifan/go-mall/admin/internal/construct"
	"net/http"
)

type JWTAuthTokenMiddleware struct {
	auth config.Auth
}
type MyCustomClaims struct {
}

func NewJWTAuthTokenMiddleware(auth config.Auth) *JWTAuthTokenMiddleware {
	return &JWTAuthTokenMiddleware{
		auth: auth,
	}
}

func (m *JWTAuthTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization token must be present", http.StatusUnauthorized)
			return
		}
		type MyCustomClaims struct {
			Foo string `json:"foo"`
			Sub string `json:"sub"`
			UID int    `json:"uid"`
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.auth.JwtSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if Claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			r = r.WithContext(context.WithValue(
				r.Context(),
				construct.ClaimsUserContextKey{},
				Claims["uid"]),
			)
			r = r.WithContext(context.WithValue(
				r.Context(),
				construct.ClaimsUIDContextKey{},
				Claims["sub"]),
			)
		}
		next(w, r)
	}
}
