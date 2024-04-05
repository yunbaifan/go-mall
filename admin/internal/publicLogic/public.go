package publicLogic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yunbaifan/go-mall/admin/internal/construct"
	"time"
)

func JwtToken(_ context.Context, user *construct.Claims) (string, error) {
	claims := make(jwt.MapClaims)
	// 设置过期时间
	now := time.Now()
	claims["created"] = time.Now()
	claims["sub"] = user.Sub
	// 过期时间
	claims["exp"] = now.Add(user.Expire).Unix()
	claims["uid"] = user.UID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(user.JwtSecret))
}
