package model

import (
	"gorm.io/gorm"
)

var _ UmsMemberProductCategoryRelationModel = (*customUmsMemberProductCategoryRelationModel)(nil)

type (
	// UmsMemberProductCategoryRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberProductCategoryRelationModel.
	UmsMemberProductCategoryRelationModel interface {
		umsMemberProductCategoryRelationModel
	}

	customUmsMemberProductCategoryRelationModel struct {
		*defaultUmsMemberProductCategoryRelationModel
	}

	UmsMemberProductCategoryRelationRelation struct {
		//
	}
)

// NewUmsMemberProductCategoryRelationModel returns a model for the database table.
func NewUmsMemberProductCategoryRelationModel(db *gorm.DB) UmsMemberProductCategoryRelationModel {
	return &customUmsMemberProductCategoryRelationModel{
		defaultUmsMemberProductCategoryRelationModel: newUmsMemberProductCategoryRelationModel(db),
	}
}
