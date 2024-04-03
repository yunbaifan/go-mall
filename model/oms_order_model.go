package model

import (
	"gorm.io/gorm"
)

var _ OmsOrderModel = (*customOmsOrderModel)(nil)

type (
	// OmsOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderModel.
	OmsOrderModel interface {
		omsOrderModel
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
