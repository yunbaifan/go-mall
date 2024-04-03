package model

import (
	"gorm.io/gorm"
)

var _ OmsOrderReturnReasonModel = (*customOmsOrderReturnReasonModel)(nil)

type (
	// OmsOrderReturnReasonModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderReturnReasonModel.
	OmsOrderReturnReasonModel interface {
		omsOrderReturnReasonModel
	}

	customOmsOrderReturnReasonModel struct {
		*defaultOmsOrderReturnReasonModel
	}

	OmsOrderReturnReasonRelation struct {
		//
	}
)

// NewOmsOrderReturnReasonModel returns a model for the database table.
func NewOmsOrderReturnReasonModel(db *gorm.DB) OmsOrderReturnReasonModel {
	return &customOmsOrderReturnReasonModel{
		defaultOmsOrderReturnReasonModel: newOmsOrderReturnReasonModel(db),
	}
}
