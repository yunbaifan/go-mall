package model

import (
	"gorm.io/gorm"
)

var _ UmsAdminLoginLogModel = (*customUmsAdminLoginLogModel)(nil)

type (
	// UmsAdminLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsAdminLoginLogModel.
	UmsAdminLoginLogModel interface {
		umsAdminLoginLogModel
	}

	customUmsAdminLoginLogModel struct {
		*defaultUmsAdminLoginLogModel
	}

	UmsAdminLoginLogRelation struct {
		//
	}
)

// NewUmsAdminLoginLogModel returns a model for the database table.
func NewUmsAdminLoginLogModel(db *gorm.DB) UmsAdminLoginLogModel {
	return &customUmsAdminLoginLogModel{
		defaultUmsAdminLoginLogModel: newUmsAdminLoginLogModel(db),
	}
}
