package admin

import (
	"context"
	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"
	"github.com/yunbaifan/go-mall/lib/xcode"
	"golang.org/x/crypto/bcrypt"

	"database/sql"
	"github.com/yunbaifan/go-mall/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type AdminRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminRegisterLogic {
	return &AdminRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminRegisterLogic) AdminRegister(req *types.RegisterRequest) (resp *types.RegisterRequest, err error) {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Errorf("adminRegisterLogic.AdminRegister.bcrypt.GenerateFromPassword err:%v", err)
		return nil, l.svcCtx.ResponseInter.Error(xcode.ErrInvalidOldPassword)
	}
	if err := l.svcCtx.UmsAdminModel.Insert(l.ctx, &model.UmsAdmin{
		Username: sql.NullString{String: req.Username, Valid: true},
		Password: sql.NullString{String: string(password), Valid: true},
		Email:    sql.NullString{String: req.Email, Valid: true},
		Note:     sql.NullString{String: req.Note, Valid: true},
		NickName: sql.NullString{String: req.NickName, Valid: true},
		Status:   req.Status,
	}); err != nil {
		l.Logger.Errorf("adminRegisterLogic.AdminRegister.Insert err:%v", err)
		return nil, l.svcCtx.ResponseInter.Error(xcode.ErrDataInsertFailed)
	}
	return req, nil
}
