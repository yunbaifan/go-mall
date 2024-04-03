package admin

import (
	"context"
	"encoding/json"
	"github.com/yunbaifan/go-mall/lib/xorm"
	"github.com/yunbaifan/go-mall/model"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

var (
	_db *gorm.DB
)

func TestMain(m *testing.M) {
	var (
		err error
	)
	_db, err = xorm.ConnectMysql(xorm.DatabaseConf{
		Source:        "",
		MaxIdleConns:  10,
		MaxOpenConns:  10,
		SlowThreshold: 100,
		LogLevel:      logger.LogLevel(4),
		Colorful:      false,
	})
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestFindByAdminID(t *testing.T) {
	var (
		adminID int64 = 3
	)
	ctx := context.Background()
	admin, err := model.NewUmsAdminRoleRelationModel(_db).FindByAdminID(ctx, adminID, "UmsAdminRoleRelationRelation.UmsRoles")
	if err != nil {
		t.Error(err)
	}
	roles := make([]string, 0, 0)
	roles = append(roles, admin.UmsAdminRoleRelationRelation.UmsRoles.Name.String)

	data, err := model.NewUmsRoleMenuRelationModel(_db).FindByRoleID(ctx, admin.RoleId.Int64, "UmsRoleMenuRelationRelation.UmsMenus")
	jsonData, _ := json.Marshal(data)
	t.Logf("data: %s", string(jsonData))
	t.Logf("admin: %v", roles)
}

//func Test
