package model

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ UmsAdminModel = (*customUmsAdminModel)(nil)

type (
	// UmsAdminModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsAdminModel.
	//
	UmsAdminModel interface {
		umsAdminModel
		umsAdminInter
	}

	// 另外一个接口
	umsAdminInter interface {
		FindOneByUserName(ctx context.Context, username string) (*UmsAdmin, error)
		UpdateSubQuery(ctx context.Context, id int64) (rowsAffected int64, err error)
	}

	customUmsAdminModel struct {
		*defaultUmsAdminModel
	}

	UmsAdminRelation struct {
		//
	}
)

// NewUmsAdminModel returns a model for the database table.
func NewUmsAdminModel(db *gorm.DB) UmsAdminModel {
	return &customUmsAdminModel{
		defaultUmsAdminModel: newUmsAdminModel(db),
	}
}

func (d *defaultUmsAdminModel) FindOneByUserName(ctx context.Context, username string) (*UmsAdmin, error) {
	var mm UmsAdmin
	err := d.OrmSession(ctx).Where("username = ?", username).First(&mm).Error
	switch err {
	case nil:
		return &mm, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

type SuQuery struct {
	Id int64
}

// GormValue 可以设置子查询
func (s SuQuery) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "SELECT `status` FROM (SELECT `status` FROM ums_admin WHERE id = ?) AS subquery",
		Vars: []interface{}{s.Id},
	}
}

func (d *defaultUmsAdminModel) UpdateSubQuery(ctx context.Context, id int64) (int64, error) {
	db := d.OrmSession(ctx)
	res := db.Model(&UmsAdmin{}).Where("id = ?", id).
		Update("status", gorm.Expr("1 ^ (?)", SuQuery{Id: id}))
	return res.RowsAffected, res.Error
}
