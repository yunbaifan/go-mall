package model

import (
	"gorm.io/gorm"
)

var _ SmsCouponProductCategoryRelationModel = (*customSmsCouponProductCategoryRelationModel)(nil)

type (
	// SmsCouponProductCategoryRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsCouponProductCategoryRelationModel.
	SmsCouponProductCategoryRelationModel interface {
		smsCouponProductCategoryRelationModel
	}

	customSmsCouponProductCategoryRelationModel struct {
		*defaultSmsCouponProductCategoryRelationModel
	}

	SmsCouponProductCategoryRelationRelation struct {
		//
	}
)

// NewSmsCouponProductCategoryRelationModel returns a model for the database table.
func NewSmsCouponProductCategoryRelationModel(db *gorm.DB) SmsCouponProductCategoryRelationModel {
	return &customSmsCouponProductCategoryRelationModel{
		defaultSmsCouponProductCategoryRelationModel: newSmsCouponProductCategoryRelationModel(db),
	}
}
