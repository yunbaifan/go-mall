package admin

import (
	"context"
	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"
	"github.com/yunbaifan/go-mall/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminInfoLogic {
	return &AdminInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminInfoLogic) AdminInfo() (resp *types.AdminInfo, err error) {
	userID := l.ctx.Value("userID").(int64)
	// 查询用户信息
	var (
		res *model.UmsAdminRoleRelation
	)
	if res, err = l.svcCtx.UmsAdminRoleRelationModel.FindByAdminID(
		l.ctx, userID,
		"UmsAdminRoleRelationRelation.UmsRoles",
	); err != nil {
		l.Logger.Errorf("AdminLogic.Admin",
			logx.Field("err", err),
		)
		return
	}
	roles := make([]string, 0, 0)
	roles = append(roles, res.UmsAdminRoleRelationRelation.UmsRoles.Name.String)
	// 查询用户菜单
	var (
		data []model.UmsRoleMenuRelation
	)

	if data, err = l.svcCtx.UmsRoleMenuRelationModel.FindByRoleID(
		l.ctx, res.RoleId.Int64,
		"UmsRoleMenuRelationRelation.UmsMenus",
	); err != nil {
		l.Logger.Errorf("AdminInfo.FindByRoleID",
			logx.Field("err", err),
		)
		return
	}
	menus := make([]types.Menus, 0, len(data))
	for k := range data {
		val := data[k]
		menus = append(menus, types.Menus{
			Name:       val.UmsRoleMenuRelationRelation.UmsMenus.Name.String,
			ParentId:   int(val.UmsRoleMenuRelationRelation.UmsMenus.ParentId.Int64),
			CreateTime: val.UmsRoleMenuRelationRelation.UmsMenus.CreateTime.Time.Format("2006-01-02 15:04:05"),
			Level:      int(val.UmsRoleMenuRelationRelation.UmsMenus.Level.Int64),
			Sort:       int(val.UmsRoleMenuRelationRelation.UmsMenus.Sort.Int64),
			Icon:       val.UmsRoleMenuRelationRelation.UmsMenus.Icon.String,
			Hidden:     val.UmsRoleMenuRelationRelation.UmsMenus.Hidden.Int64 == 1,
		})
	}
	resp = &types.AdminInfo{
		Icon:  "",
		Roles: roles,
		Menus: menus,
	}
	// 释放资源
	menus = nil
	roles = nil
	data = nil
	return
}
