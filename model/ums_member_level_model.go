package model

import (
	"gorm.io/gorm"
)

var _ UmsMemberLevelModel = (*customUmsMemberLevelModel)(nil)

type (
	// UmsMemberLevelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberLevelModel.
	//
	UmsMemberLevelModel interface {
		umsMemberLevelModel
		umsMemberLevelInter
	}

	// 另外一个接口
	umsMemberLevelInter interface {
		// add custom methods here
	}

	customUmsMemberLevelModel struct {
		*defaultUmsMemberLevelModel
	}

	UmsMemberLevelRelation struct {
		//
	}
)

// NewUmsMemberLevelModel returns a model for the database table.
func NewUmsMemberLevelModel(db *gorm.DB) UmsMemberLevelModel {
	return &customUmsMemberLevelModel{
		defaultUmsMemberLevelModel: newUmsMemberLevelModel(db),
	}
}
