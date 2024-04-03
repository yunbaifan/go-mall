package model

import (
	"gorm.io/gorm"
)

var _ OmsOrderItemModel = (*customOmsOrderItemModel)(nil)

type (
	// OmsOrderItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderItemModel.
	OmsOrderItemModel interface {
		omsOrderItemModel
	}

	customOmsOrderItemModel struct {
		*defaultOmsOrderItemModel
	}

	OmsOrderItemRelation struct {
		//
	}
)

// NewOmsOrderItemModel returns a model for the database table.
func NewOmsOrderItemModel(db *gorm.DB) OmsOrderItemModel {
	return &customOmsOrderItemModel{
		defaultOmsOrderItemModel: newOmsOrderItemModel(db),
	}
}
