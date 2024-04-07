package admin

import (
	"context"
	"database/sql"
	"github.com/yunbaifan/go-mall/lib/xcode"
	"github.com/yunbaifan/go-mall/model"
	"strconv"
	"strings"

	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminRoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminRoleUpdateLogic {
	return &AdminRoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminRoleUpdateLogic) AdminRoleUpdate(req *types.AdminRoleUpdateRequest) (resp *types.AdminCountResponse, err error) {
	//先根据 adminId 删除所有关联关系
	if err = l.svcCtx.UmsAdminRoleRelationModel.DeleteByAdminID(l.ctx, req.AdminId); err != nil {
		l.Logger.Errorf("AdminRoleUpdateLogic.UmsAdminRoleRelationModel.DeleteByAdminID",
			logx.Field("err", err),
		)
		return nil, l.svcCtx.ResponseInter.Error(xcode.ErrDataDeleteFailed)
	}
	roleIds := strings.Split(req.RoleIDs, ",")
	list := make([]*model.UmsAdminRoleRelation, 0, len(roleIds))
	for k := range roleIds {
		var (
			roleId int64
		)
		val := roleIds[k]
		roleId, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			l.Logger.Errorf("AdminRoleUpdateLogic.AdminRoleUpdate.Range",
				logx.Field("err", err),
			)
			continue
		}
		list = append(list, &model.UmsAdminRoleRelation{
			AdminId: sql.NullInt64{Int64: req.AdminId},
			RoleId:  sql.NullInt64{Int64: roleId},
		})
	}
	var (
		rowsAffected int64
	)
	if rowsAffected, err = l.svcCtx.UmsAdminRoleRelationModel.BatchInsert(l.ctx, list); err != nil {
		l.Logger.Errorf("AdminRoleUpdateLogic.UmsAdminRoleRelationModel.InsertMulti",
			logx.Field("err", err),
		)
		return nil, l.svcCtx.ResponseInter.Error(xcode.ErrDataInsertFailed)
	}
	// 然后重新插入
	return &types.AdminCountResponse{
		Count: rowsAffected,
	}, nil
}
