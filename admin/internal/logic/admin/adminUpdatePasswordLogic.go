package admin

import (
	"context"
	"database/sql"
	"github.com/yunbaifan/go-mall/lib/xcode"
	"github.com/yunbaifan/go-mall/model"
	"golang.org/x/crypto/bcrypt"

	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUpdatePasswordLogic {
	return &AdminUpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminUpdatePasswordLogic) AdminUpdatePassword(req *types.UpdatePasswordRequest) (resp *types.AdminCountResponse, err error) {
	//查询用户密码
	admin, err := l.svcCtx.UmsAdminModel.FindOne(l.ctx, req.UserID)
	if err != nil {
		l.Logger.Errorf("AdminUpdatePasswordLogic.AdminUpdatePassword.FindOne",
			logx.Field("err", err),
		)
		return nil, l.svcCtx.ResponseInter.Error(xcode.ErrDataNotFoundFailed)
	}
	//验证密码
	if err = bcrypt.CompareHashAndPassword([]byte(admin.Password.String), []byte(req.OldPassword)); err != nil {
		l.Logger.Errorf("AdminUpdatePasswordLogic.AdminUpdatePassword.CompareHashAndPassword",
			logx.Field("err", err),
		)
		return nil, l.svcCtx.ResponseInter.Error(xcode.ErrInvalidOldPassword)
	}
	//生成新密码
	password, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Errorf("AdminUpdatePasswordLogic.AdminUpdatePassword.GenerateFromPassword",
			logx.Field("err", err),
		)
		return nil, l.svcCtx.ResponseInter.Error(xcode.ErrThirdPartyPluginError)
	}
	//更新密码
	if err = l.svcCtx.UmsAdminModel.Update(l.ctx, &model.UmsAdmin{
		Id:       req.UserID,
		Password: sql.NullString{String: string(password), Valid: true},
	}); err != nil {
		l.Logger.Errorf("AdminUpdatePasswordLogic.AdminUpdatePassword.UpdatePassword",
			logx.Field("err", err),
		)
		return nil, l.svcCtx.ResponseInter.Error(xcode.ErrUpdateDataFailed)
	}
	return &types.AdminCountResponse{
		Count: 1,
	}, nil
}
