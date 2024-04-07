package admin

import (
	"context"
	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"
	"github.com/yunbaifan/go-mall/lib/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUpdateStatusLogic {
	return &AdminUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminUpdateStatusLogic) AdminUpdateStatus(req *types.AdminGetItemRequest) (resp *types.AdminCountResponse, err error) {
	var (
		rowsAffected int64
	)
	if rowsAffected, err = l.svcCtx.UmsAdminModel.UpdateSubQuery(l.ctx, req.UserID); err != nil {
		l.Logger.Errorf("AdminUpdateStatusLogic.AdminUpdateStatus.UpdateSubQuery",
			logx.Field("err", err),
		)
		return nil, l.svcCtx.ResponseInter.Error(xcode.ErrUpdateDataFailed)
	}
	return &types.AdminCountResponse{
		Count: rowsAffected,
	}, nil
}
