package admin

import (
	"context"
	"fmt"
	"github.com/yunbaifan/go-mall/admin/internal/construct"
	"github.com/yunbaifan/go-mall/admin/internal/publicLogic"
	"time"

	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken() (resp *types.LoginResponse, err error) {
	userID := l.ctx.Value(construct.ClaimsUIDContextKey{}).(int64)
	userName := l.ctx.Value(construct.ClaimsUserContextKey{}).(string)
	token, err := publicLogic.JwtToken(l.ctx, &construct.Claims{
		UID:       userID,
		Sub:       userName,
		Expire:    time.Duration(l.svcCtx.Config.Auth.Expire) * time.Second,
		JwtSecret: l.svcCtx.Config.Auth.JwtSecret,
	})
	if err != nil {
		l.Logger.Errorf("RefreshTokenLogic.RefreshToken.jwtToken",
			logx.Field("err", err),
		)
		return nil, fmt.Errorf("failed to create token %v", err)
	}
	return &types.LoginResponse{
		Token:     token,
		TokenHead: "Bearer ",
	}, nil
}
