package model

import (
	"gorm.io/gorm"
)

var _ PmsProductModel = (*customPmsProductModel)(nil)

type (
	// PmsProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductModel.
	//
	PmsProductModel interface {
		pmsProductModel
		pmsProductInter
	}

	// 另外一个接口
	pmsProductInter interface {
		// add custom methods here
	}

	customPmsProductModel struct {
		*defaultPmsProductModel
	}

	PmsProductRelation struct {
		//
	}
)

// NewPmsProductModel returns a model for the database table.
func NewPmsProductModel(db *gorm.DB) PmsProductModel {
	return &customPmsProductModel{
		defaultPmsProductModel: newPmsProductModel(db),
	}
}
