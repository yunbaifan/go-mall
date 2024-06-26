package model

import (
	"gorm.io/gorm"
)

var _ PmsProductCategoryModel = (*customPmsProductCategoryModel)(nil)

type (
	// PmsProductCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductCategoryModel.
	//
	PmsProductCategoryModel interface {
		pmsProductCategoryModel
		pmsProductCategoryInter
	}

	// 另外一个接口
	pmsProductCategoryInter interface {
		// add custom methods here
	}

	customPmsProductCategoryModel struct {
		*defaultPmsProductCategoryModel
	}

	PmsProductCategoryRelation struct {
		//
	}
)

// NewPmsProductCategoryModel returns a model for the database table.
func NewPmsProductCategoryModel(db *gorm.DB) PmsProductCategoryModel {
	return &customPmsProductCategoryModel{
		defaultPmsProductCategoryModel: newPmsProductCategoryModel(db),
	}
}
