package model

import (
	"gorm.io/gorm"
)

var _ UmsAdminPermissionRelationModel = (*customUmsAdminPermissionRelationModel)(nil)

type (
	// UmsAdminPermissionRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsAdminPermissionRelationModel.
	UmsAdminPermissionRelationModel interface {
		umsAdminPermissionRelationModel
	}

	customUmsAdminPermissionRelationModel struct {
		*defaultUmsAdminPermissionRelationModel
	}

	UmsAdminPermissionRelationRelation struct {
		//
	}
)

// NewUmsAdminPermissionRelationModel returns a model for the database table.
func NewUmsAdminPermissionRelationModel(db *gorm.DB) UmsAdminPermissionRelationModel {
	return &customUmsAdminPermissionRelationModel{
		defaultUmsAdminPermissionRelationModel: newUmsAdminPermissionRelationModel(db),
	}
}
