package model

import (
	"gorm.io/gorm"
)

var _ UmsMemberTagModel = (*customUmsMemberTagModel)(nil)

type (
	// UmsMemberTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberTagModel.
	UmsMemberTagModel interface {
		umsMemberTagModel
	}

	customUmsMemberTagModel struct {
		*defaultUmsMemberTagModel
	}

	UmsMemberTagRelation struct {
		//
	}
)

// NewUmsMemberTagModel returns a model for the database table.
func NewUmsMemberTagModel(db *gorm.DB) UmsMemberTagModel {
	return &customUmsMemberTagModel{
		defaultUmsMemberTagModel: newUmsMemberTagModel(db),
	}
}
