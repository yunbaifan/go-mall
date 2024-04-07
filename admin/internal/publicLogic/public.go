package publicLogic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yunbaifan/go-mall/admin/internal/construct"
	"github.com/yunbaifan/go-mall/lib/xcode"
	"golang.org/x/text/language"
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

type ResponseInter interface {
	Error(code xcode.ErrorCode) error
}
type responseInter struct {
	tag language.Tag
}

func NewResponseInter(tag language.Tag) ResponseInter {
	return &responseInter{tag: tag}
}

func (r *responseInter) Error(code xcode.ErrorCode) error {
	return xcode.NewResponse(code, r.tag)
}
