package admin

import (
	"context"
	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"
	"github.com/yunbaifan/go-mall/lib/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminGetItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetItemLogic {
	return &AdminGetItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetItemLogic) AdminGetItem(req *types.AdminGetItemRequest) (resp *types.AdminResponse, err error) {
	res, err := l.svcCtx.UmsAdminModel.FindOne(l.ctx, req.UserID)
	if err != nil {
		l.Logger.Errorf("AdminGetItemLogic.AdminGetItem.FindOne",
			logx.Field("err", err),
		)
		return nil, l.svcCtx.ResponseInter.Error(xcode.ErrDataNotFoundFailed)
	}
	return &types.AdminResponse{
		ID:         res.Id,
		Username:   res.Username.String,
		Password:   res.Password.String,
		Icon:       res.Icon.String,
		Email:      res.Email.String,
		NickName:   res.NickName.String,
		Note:       res.Note.String,
		CreateTime: res.CreateTime.Time.Format("2006-01-02 15:04:05"),
		LoginTime:  res.LoginTime.Time.Format("2006-01-02 15:04:05"),
		Status:     res.Status,
	}, nil
}
