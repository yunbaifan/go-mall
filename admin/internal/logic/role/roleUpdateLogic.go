package role

import (
	"context"

	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleUpdateLogic) RoleUpdate(req *types.RoleRequest) (resp *types.RoleCountResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
