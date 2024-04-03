package model

import (
	"gorm.io/gorm"
)

var _ SmsCouponModel = (*customSmsCouponModel)(nil)

type (
	// SmsCouponModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsCouponModel.
	SmsCouponModel interface {
		smsCouponModel
	}

	customSmsCouponModel struct {
		*defaultSmsCouponModel
	}

	SmsCouponRelation struct {
		//
	}
)

// NewSmsCouponModel returns a model for the database table.
func NewSmsCouponModel(db *gorm.DB) SmsCouponModel {
	return &customSmsCouponModel{
		defaultSmsCouponModel: newSmsCouponModel(db),
	}
}
