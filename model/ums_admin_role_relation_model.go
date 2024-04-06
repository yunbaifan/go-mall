package model

import (
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
		// add custom methods here
	}

	customUmsAdminRoleRelationModel struct {
		*defaultUmsAdminRoleRelationModel
	}

	UmsAdminRoleRelationRelation struct {
		//
	}
)

// NewUmsAdminRoleRelationModel returns a model for the database table.
func NewUmsAdminRoleRelationModel(db *gorm.DB) UmsAdminRoleRelationModel {
	return &customUmsAdminRoleRelationModel{
		defaultUmsAdminRoleRelationModel: newUmsAdminRoleRelationModel(db),
	}
}
