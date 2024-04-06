package model

import (
	"gorm.io/gorm"
)

var _ PmsProductAttributeCategoryModel = (*customPmsProductAttributeCategoryModel)(nil)

type (
	// PmsProductAttributeCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductAttributeCategoryModel.
	//
	PmsProductAttributeCategoryModel interface {
		pmsProductAttributeCategoryModel
		pmsProductAttributeCategoryInter
	}

	// 另外一个接口
	pmsProductAttributeCategoryInter interface {
		// add custom methods here
	}

	customPmsProductAttributeCategoryModel struct {
		*defaultPmsProductAttributeCategoryModel
	}

	PmsProductAttributeCategoryRelation struct {
		//
	}
)

// NewPmsProductAttributeCategoryModel returns a model for the database table.
func NewPmsProductAttributeCategoryModel(db *gorm.DB) PmsProductAttributeCategoryModel {
	return &customPmsProductAttributeCategoryModel{
		defaultPmsProductAttributeCategoryModel: newPmsProductAttributeCategoryModel(db),
	}
}
