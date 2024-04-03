package model

import (
	"gorm.io/gorm"
)

var _ SmsCouponProductRelationModel = (*customSmsCouponProductRelationModel)(nil)

type (
	// SmsCouponProductRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsCouponProductRelationModel.
	SmsCouponProductRelationModel interface {
		smsCouponProductRelationModel
	}

	customSmsCouponProductRelationModel struct {
		*defaultSmsCouponProductRelationModel
	}

	SmsCouponProductRelationRelation struct {
		//
	}
)

// NewSmsCouponProductRelationModel returns a model for the database table.
func NewSmsCouponProductRelationModel(db *gorm.DB) SmsCouponProductRelationModel {
	return &customSmsCouponProductRelationModel{
		defaultSmsCouponProductRelationModel: newSmsCouponProductRelationModel(db),
	}
}
