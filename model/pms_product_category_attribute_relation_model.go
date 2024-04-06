package model

import (
	"gorm.io/gorm"
)

var _ PmsProductCategoryAttributeRelationModel = (*customPmsProductCategoryAttributeRelationModel)(nil)

type (
	// PmsProductCategoryAttributeRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductCategoryAttributeRelationModel.
	//
	PmsProductCategoryAttributeRelationModel interface {
		pmsProductCategoryAttributeRelationModel
		pmsProductCategoryAttributeRelationInter
	}

	// 另外一个接口
	pmsProductCategoryAttributeRelationInter interface {
		// add custom methods here
	}

	customPmsProductCategoryAttributeRelationModel struct {
		*defaultPmsProductCategoryAttributeRelationModel
	}

	PmsProductCategoryAttributeRelationRelation struct {
		//
	}
)

// NewPmsProductCategoryAttributeRelationModel returns a model for the database table.
func NewPmsProductCategoryAttributeRelationModel(db *gorm.DB) PmsProductCategoryAttributeRelationModel {
	return &customPmsProductCategoryAttributeRelationModel{
		defaultPmsProductCategoryAttributeRelationModel: newPmsProductCategoryAttributeRelationModel(db),
	}
}
