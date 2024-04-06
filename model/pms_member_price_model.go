package model

import (
	"gorm.io/gorm"
)

var _ PmsMemberPriceModel = (*customPmsMemberPriceModel)(nil)

type (
	// PmsMemberPriceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsMemberPriceModel.
	//
	PmsMemberPriceModel interface {
		pmsMemberPriceModel
		pmsMemberPriceInter
	}

	// 另外一个接口
	pmsMemberPriceInter interface {
		// add custom methods here
	}

	customPmsMemberPriceModel struct {
		*defaultPmsMemberPriceModel
	}

	PmsMemberPriceRelation struct {
		//
	}
)

// NewPmsMemberPriceModel returns a model for the database table.
func NewPmsMemberPriceModel(db *gorm.DB) PmsMemberPriceModel {
	return &customPmsMemberPriceModel{
		defaultPmsMemberPriceModel: newPmsMemberPriceModel(db),
	}
}
