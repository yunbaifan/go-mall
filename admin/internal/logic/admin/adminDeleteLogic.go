package admin

import (
	"context"
	"github.com/yunbaifan/go-mall/lib/xcode"

	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminDeleteLogic {
	return &AdminDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminDeleteLogic) AdminDelete(req *types.AdminGetItemRequest) (resp *types.AdminCountResponse, err error) {
	if err = l.svcCtx.UmsAdminModel.Delete(l.ctx, req.UserID); err != nil {
		l.Logger.Errorf("AdminGetItemLogic.AdminDelete.Delete",
			logx.Field("err", err),
		)
		return nil, l.svcCtx.ResponseInter.Error(xcode.ErrDataDeleteFailed)
	}
	return &types.AdminCountResponse{
		Count: 1,
	}, nil
}
