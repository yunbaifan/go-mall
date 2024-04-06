package model

import (
	"gorm.io/gorm"
)

var _ PmsBrandModel = (*customPmsBrandModel)(nil)

type (
	// PmsBrandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsBrandModel.
	//
	PmsBrandModel interface {
		pmsBrandModel
		pmsBrandInter
	}

	// 另外一个接口
	pmsBrandInter interface {
		// add custom methods here
	}

	customPmsBrandModel struct {
		*defaultPmsBrandModel
	}

	PmsBrandRelation struct {
		//
	}
)

// NewPmsBrandModel returns a model for the database table.
func NewPmsBrandModel(db *gorm.DB) PmsBrandModel {
	return &customPmsBrandModel{
		defaultPmsBrandModel: newPmsBrandModel(db),
	}
}
