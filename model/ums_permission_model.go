package model

import (
	"gorm.io/gorm"
)

var _ UmsPermissionModel = (*customUmsPermissionModel)(nil)

type (
	// UmsPermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsPermissionModel.
	UmsPermissionModel interface {
		umsPermissionModel
	}

	customUmsPermissionModel struct {
		*defaultUmsPermissionModel
	}

	UmsPermissionRelation struct {
		//
	}
)

// NewUmsPermissionModel returns a model for the database table.
func NewUmsPermissionModel(db *gorm.DB) UmsPermissionModel {
	return &customUmsPermissionModel{
		defaultUmsPermissionModel: newUmsPermissionModel(db),
	}
}
