package model

import (
	"gorm.io/gorm"
)

var _ SmsCouponHistoryModel = (*customSmsCouponHistoryModel)(nil)

type (
	// SmsCouponHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsCouponHistoryModel.
	SmsCouponHistoryModel interface {
		smsCouponHistoryModel
	}

	customSmsCouponHistoryModel struct {
		*defaultSmsCouponHistoryModel
	}

	SmsCouponHistoryRelation struct {
		//
	}
)

// NewSmsCouponHistoryModel returns a model for the database table.
func NewSmsCouponHistoryModel(db *gorm.DB) SmsCouponHistoryModel {
	return &customSmsCouponHistoryModel{
		defaultSmsCouponHistoryModel: newSmsCouponHistoryModel(db),
	}
}
