package model

import (
	"gorm.io/gorm"
)

var _ OmsOrderReturnApplyModel = (*customOmsOrderReturnApplyModel)(nil)

type (
	// OmsOrderReturnApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderReturnApplyModel.
	//
	OmsOrderReturnApplyModel interface {
		omsOrderReturnApplyModel
		omsOrderReturnApplyInter
	}

	// 另外一个接口
	omsOrderReturnApplyInter interface {
		// add custom methods here
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
