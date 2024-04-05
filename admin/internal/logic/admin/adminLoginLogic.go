package admin

import (
	"context"
	"errors"
	"fmt"
	"github.com/yunbaifan/go-mall/admin/internal/construct"
	"github.com/yunbaifan/go-mall/admin/internal/publicLogic"
	"github.com/yunbaifan/go-mall/model"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("username or password cannot be empty")
	}
	// 查询用户信息
	var (
		res *model.UmsAdmin
	)
	if res, err = l.svcCtx.UmsAdminModel.FindOneByUserName(l.ctx, req.Username); err != nil {
		l.Logger.Errorf("AdminLogic.AdminLogin.FindOneByUserName",
			logx.Field("err", err),
		)
		return
	}
	// 判断用户密码是否正确
	if err = bcrypt.CompareHashAndPassword(
		[]byte(res.Password.String),
		[]byte(req.Password),
	); err != nil {
		return nil, fmt.Errorf("password doesn't match %s", req.Username)
	}
	// 判断用户是否被禁用
	if res.Status != 1 {
		return nil, fmt.Errorf("user is disabled %s", req.Username)
	}
	token, err := publicLogic.JwtToken(l.ctx, &construct.Claims{
		UID:       res.Id,
		Sub:       res.Username.String,
		Expire:    time.Duration(l.svcCtx.Config.Auth.Expire) * time.Second,
		JwtSecret: l.svcCtx.Config.Auth.JwtSecret,
	})
	if err != nil {
		l.Logger.Errorf("AdminLogic.AdminLogin.jwtToken",
			logx.Field("err", err),
		)
		return nil, fmt.Errorf("failed to create token %v", err)
	}
	return &types.LoginResponse{
		Token:     token,
		TokenHead: "Bearer ",
	}, nil
}
