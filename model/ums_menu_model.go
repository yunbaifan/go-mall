package model

import (
	"gorm.io/gorm"
)

var _ UmsMenuModel = (*customUmsMenuModel)(nil)

type (
	// UmsMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMenuModel.
	UmsMenuModel interface {
		umsMenuModel
	}

	customUmsMenuModel struct {
		*defaultUmsMenuModel
	}

	UmsMenuRelation struct {
		//
	}
)

// NewUmsMenuModel returns a model for the database table.
func NewUmsMenuModel(db *gorm.DB) UmsMenuModel {
	return &customUmsMenuModel{
		defaultUmsMenuModel: newUmsMenuModel(db),
	}
}
