package svc

import (
	"github.com/yunbaifan/go-mall/admin/internal/config"
	"github.com/yunbaifan/go-mall/admin/internal/middleware"
	"github.com/yunbaifan/go-mall/lib/xorm"
	"github.com/yunbaifan/go-mall/model"
	"github.com/zeromicro/go-zero/rest"
	"log"
)

type ServiceContext struct {
	Config                    config.Config
	UmsAdminModel             model.UmsAdminModel
	UmsRoleModel              model.UmsRoleModel
	UmsAdminRoleRelationModel model.UmsAdminRoleRelationModel
	UmsMenuModel              model.UmsMenuModel
	UmsRoleMenuRelationModel  model.UmsRoleMenuRelationModel
	JWTAuthToken              rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := xorm.ConnectMysql(c.MySQL)
	if err != nil {
		log.Fatalf("error: xorm.ConnectMysql: %s", err.Error())
	}
	return &ServiceContext{
		Config: c,
		//后台用户表
		UmsAdminModel: model.NewUmsAdminModel(db),
		//后台用户角色表
		UmsRoleModel: model.NewUmsRoleModel(db),
		//后台用户和角色关系表
		UmsAdminRoleRelationModel: model.NewUmsAdminRoleRelationModel(db),
		UmsMenuModel:              model.NewUmsMenuModel(db),
		UmsRoleMenuRelationModel:  model.NewUmsRoleMenuRelationModel(db),
		JWTAuthToken:              middleware.NewJWTAuthTokenMiddleware(c.Auth).Handle,
	}
}
