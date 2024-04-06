package model

import (
	"gorm.io/gorm"
)

var _ UmsMemberLoginLogModel = (*customUmsMemberLoginLogModel)(nil)

type (
	// UmsMemberLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberLoginLogModel.
	//
	UmsMemberLoginLogModel interface {
		umsMemberLoginLogModel
		umsMemberLoginLogInter
	}

	// 另外一个接口
	umsMemberLoginLogInter interface {
		// add custom methods here
	}

	customUmsMemberLoginLogModel struct {
		*defaultUmsMemberLoginLogModel
	}

	UmsMemberLoginLogRelation struct {
		//
	}
)

// NewUmsMemberLoginLogModel returns a model for the database table.
func NewUmsMemberLoginLogModel(db *gorm.DB) UmsMemberLoginLogModel {
	return &customUmsMemberLoginLogModel{
		defaultUmsMemberLoginLogModel: newUmsMemberLoginLogModel(db),
	}
}
