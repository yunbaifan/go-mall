package model

import (
	"context"
	"gorm.io/gorm"
)

var _ UmsAdminRoleRelationModel = (*customUmsAdminRoleRelationModel)(nil)

type (
	// UmsAdminRoleRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsAdminRoleRelationModel.
	UmsAdminRoleRelationModel interface {
		umsAdminRoleRelationModel
		umsAdminRoleRelationInter
	}

	umsAdminRoleRelationInter interface {
		FindByAdminID(ctx context.Context, id int64, preloadList ...string) (*UmsAdminRoleRelation, error)
	}

	customUmsAdminRoleRelationModel struct {
		*defaultUmsAdminRoleRelationModel
	}

	UmsAdminRoleRelationRelation struct {
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
