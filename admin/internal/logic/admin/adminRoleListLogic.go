package admin

import (
	"context"

	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminRoleListLogic {
	return &AdminRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminRoleListLogic) AdminRoleList(req *types.AdminRoleListRequest) (resp []types.AdminRoleListResponse, err error) {
	res, err := l.svcCtx.UmsRoleModel.FindOne(l.ctx, req.AdminId)
	if err != nil {
		l.Logger.Errorf("AdminRoleListLogic.AdminRoleList.FindOne",
			logx.Field("err", err),
		)
		return
	}
	return []types.AdminRoleListResponse{
		{
			Id:          res.Id,
			Name:        res.Name.String,
			Description: res.Description.String,
			AdminCount:  res.AdminCount.Int64,
			CreateTime:  res.CreateTime.Time.Format("2006-01-02 15:04:05"),
			Status:      res.Status,
			Sort:        res.Sort,
		},
	}, nil
}
