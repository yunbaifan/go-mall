package model

import (
	"gorm.io/gorm"
)

var _ UmsMemberModel = (*customUmsMemberModel)(nil)

type (
	// UmsMemberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberModel.
	//
	UmsMemberModel interface {
		umsMemberModel
		umsMemberInter
	}

	// 另外一个接口
	umsMemberInter interface {
		// add custom methods here
	}

	customUmsMemberModel struct {
		*defaultUmsMemberModel
	}

	UmsMemberRelation struct {
		//
	}
)

// NewUmsMemberModel returns a model for the database table.
func NewUmsMemberModel(db *gorm.DB) UmsMemberModel {
	return &customUmsMemberModel{
		defaultUmsMemberModel: newUmsMemberModel(db),
	}
}
