package model

import (
	"gorm.io/gorm"
)

var _ OmsOrderReturnApplyModel = (*customOmsOrderReturnApplyModel)(nil)

type (
	// OmsOrderReturnApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderReturnApplyModel.
	OmsOrderReturnApplyModel interface {
		omsOrderReturnApplyModel
	}

	customOmsOrderReturnApplyModel struct {
		*defaultOmsOrderReturnApplyModel
	}

	OmsOrderReturnApplyRelation struct {
		//
	}
)

// NewOmsOrderReturnApplyModel returns a model for the database table.
func NewOmsOrderReturnApplyModel(db *gorm.DB) OmsOrderReturnApplyModel {
	return &customOmsOrderReturnApplyModel{
		defaultOmsOrderReturnApplyModel: newOmsOrderReturnApplyModel(db),
	}
}
