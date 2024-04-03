package model

import (
	"gorm.io/gorm"
)

var _ UmsMemberTaskModel = (*customUmsMemberTaskModel)(nil)

type (
	// UmsMemberTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberTaskModel.
	UmsMemberTaskModel interface {
		umsMemberTaskModel
	}

	customUmsMemberTaskModel struct {
		*defaultUmsMemberTaskModel
	}

	UmsMemberTaskRelation struct {
		//
	}
)

// NewUmsMemberTaskModel returns a model for the database table.
func NewUmsMemberTaskModel(db *gorm.DB) UmsMemberTaskModel {
	return &customUmsMemberTaskModel{
		defaultUmsMemberTaskModel: newUmsMemberTaskModel(db),
	}
}
