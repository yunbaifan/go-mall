package model

import (
	"gorm.io/gorm"
)

var _ UmsResourceCategoryModel = (*customUmsResourceCategoryModel)(nil)

type (
	// UmsResourceCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsResourceCategoryModel.
	//
	UmsResourceCategoryModel interface {
		umsResourceCategoryModel
		umsResourceCategoryInter
	}

	// 另外一个接口
	umsResourceCategoryInter interface {
		// add custom methods here
	}

	customUmsResourceCategoryModel struct {
		*defaultUmsResourceCategoryModel
	}

	UmsResourceCategoryRelation struct {
		//
	}
)

// NewUmsResourceCategoryModel returns a model for the database table.
func NewUmsResourceCategoryModel(db *gorm.DB) UmsResourceCategoryModel {
	return &customUmsResourceCategoryModel{
		defaultUmsResourceCategoryModel: newUmsResourceCategoryModel(db),
	}
}
