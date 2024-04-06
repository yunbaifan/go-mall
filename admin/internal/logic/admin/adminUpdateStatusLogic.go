package admin

import (
	"context"

	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"

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

func (l *AdminUpdateStatusLogic) AdminUpdateStatus(req *types.AdminGetItemRequest) (resp *types.AdminRoleUpdateResponse, err error) {

	return
}
