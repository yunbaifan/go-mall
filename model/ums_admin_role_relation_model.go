package model

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

var _ UmsAdminRoleRelationModel = (*customUmsAdminRoleRelationModel)(nil)

type (
	// UmsAdminRoleRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsAdminRoleRelationModel.
	//
	UmsAdminRoleRelationModel interface {
		umsAdminRoleRelationModel
		umsAdminRoleRelationInter
	}

	// 另外一个接口
	umsAdminRoleRelationInter interface {
		FindByAdminID(ctx context.Context, id int64, preloadList ...string) (*UmsAdminRoleRelation, error)
		DeleteByAdminID(ctx context.Context, adminID int64) error
		BatchInsert(ctx context.Context, data []*UmsAdminRoleRelation) (rowsAffected int64, err error)
	}

	customUmsAdminRoleRelationModel struct {
		*defaultUmsAdminRoleRelationModel
	}

	UmsAdminRoleRelationRelation struct {
		//
		UmsRoles *UmsRole `json:"umsRole" gorm:"foreignKey:Id;references:RoleId"`
	}
)

// NewUmsAdminRoleRelationModel returns a model for the database table.
func NewUmsAdminRoleRelationModel(db *gorm.DB) UmsAdminRoleRelationModel {
	return &customUmsAdminRoleRelationModel{
		defaultUmsAdminRoleRelationModel: newUmsAdminRoleRelationModel(db),
	}
}

func (d *defaultUmsAdminRoleRelationModel) FindByAdminID(ctx context.Context, id int64, preloadList ...string) (*UmsAdminRoleRelation, error) {
	var relations *UmsAdminRoleRelation
	db := d.OrmSession(ctx)
	for _, preload := range preloadList {
		db = db.Preload(preload)
	}
	err := db.Where("admin_id = ?", id).Find(&relations).Error
	if err != nil {
		return nil, err
	}
	return relations, nil
}

func (d *defaultUmsAdminRoleRelationModel) DeleteByAdminID(ctx context.Context, adminID int64) error {
	return d.OrmSession(ctx).Delete(&UmsAdminRoleRelation{
		AdminId: sql.NullInt64{Int64: adminID},
	}).Error
}

func (d *defaultUmsAdminRoleRelationModel) BatchInsert(ctx context.Context, data []*UmsAdminRoleRelation) (rowsAffected int64, err error) {
	row := d.OrmSession(ctx).Create(data)
	if row.Error == nil && row.RowsAffected > 0 {
		return rowsAffected, nil
	}
	return 0, row.Error
}
