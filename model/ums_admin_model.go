package model

import (
	"gorm.io/gorm"
)

var _ UmsAdminModel = (*customUmsAdminModel)(nil)

type (
	// UmsAdminModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsAdminModel.
	UmsAdminModel interface {
		umsAdminModel
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
