package model

import (
	"gorm.io/gorm"
)

var _ SmsCouponModel = (*customSmsCouponModel)(nil)

type (
	// SmsCouponModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsCouponModel.
	//
	SmsCouponModel interface {
		smsCouponModel
		smsCouponInter
	}

	// 另外一个接口
	smsCouponInter interface {
		// add custom methods here
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
