package model

import (
	"context"
	"gorm.io/gorm"
)

var _ UmsRoleMenuRelationModel = (*customUmsRoleMenuRelationModel)(nil)

type (
	// UmsRoleMenuRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsRoleMenuRelationModel.
	UmsRoleMenuRelationModel interface {
		umsRoleMenuRelationModel
		umsRoleMenuRelationInter
	}

	umsRoleMenuRelationInter interface {
		FindByRoleID(ctx context.Context, id int64, preloadList ...string) ([]UmsRoleMenuRelation, error)
	}

	customUmsRoleMenuRelationModel struct {
		*defaultUmsRoleMenuRelationModel
	}

	UmsRoleMenuRelationRelation struct {
		UmsMenus *UmsMenu `json:"umsMenu" gorm:"foreignKey:Id;references:MenuId"`
	}
)

// NewUmsRoleMenuRelationModel returns a model for the database table.
func NewUmsRoleMenuRelationModel(db *gorm.DB) UmsRoleMenuRelationModel {
	return &customUmsRoleMenuRelationModel{
		defaultUmsRoleMenuRelationModel: newUmsRoleMenuRelationModel(db),
	}
}

func (d *defaultUmsRoleMenuRelationModel) FindByRoleID(ctx context.Context, id int64, preloadList ...string) ([]UmsRoleMenuRelation, error) {
	var relations []UmsRoleMenuRelation
	db := d.OrmSession(ctx)
	for _, preload := range preloadList {
		db = db.Preload(preload)
	}
	err := db.Where("role_id = ?", id).Find(&relations).Error
	if err != nil {
		return nil, err
	}
	return relations, nil
}
