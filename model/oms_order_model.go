package model

import (
	"gorm.io/gorm"
)

var _ OmsOrderModel = (*customOmsOrderModel)(nil)

type (
	// OmsOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderModel.
	//
	OmsOrderModel interface {
		omsOrderModel
		omsOrderInter
	}

	// 另外一个接口
	omsOrderInter interface {
		// add custom methods here
	}

	customOmsOrderModel struct {
		*defaultOmsOrderModel
	}

	OmsOrderRelation struct {
		//
	}
)

// NewOmsOrderModel returns a model for the database table.
func NewOmsOrderModel(db *gorm.DB) OmsOrderModel {
	return &customOmsOrderModel{
		defaultOmsOrderModel: newOmsOrderModel(db),
	}
}
