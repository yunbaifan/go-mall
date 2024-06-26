package model

import (
	"gorm.io/gorm"
)

var _ UmsRolePermissionRelationModel = (*customUmsRolePermissionRelationModel)(nil)

type (
	// UmsRolePermissionRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsRolePermissionRelationModel.
	//
	UmsRolePermissionRelationModel interface {
		umsRolePermissionRelationModel
		umsRolePermissionRelationInter
	}

	// 另外一个接口
	umsRolePermissionRelationInter interface {
		// add custom methods here
	}

	customUmsRolePermissionRelationModel struct {
		*defaultUmsRolePermissionRelationModel
	}

	UmsRolePermissionRelationRelation struct {
		//
	}
)

// NewUmsRolePermissionRelationModel returns a model for the database table.
func NewUmsRolePermissionRelationModel(db *gorm.DB) UmsRolePermissionRelationModel {
	return &customUmsRolePermissionRelationModel{
		defaultUmsRolePermissionRelationModel: newUmsRolePermissionRelationModel(db),
	}
}
