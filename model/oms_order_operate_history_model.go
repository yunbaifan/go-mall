package model

import (
	"gorm.io/gorm"
)

var _ OmsOrderOperateHistoryModel = (*customOmsOrderOperateHistoryModel)(nil)

type (
	// OmsOrderOperateHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderOperateHistoryModel.
	//
	OmsOrderOperateHistoryModel interface {
		omsOrderOperateHistoryModel
		omsOrderOperateHistoryInter
	}

	// 另外一个接口
	omsOrderOperateHistoryInter interface {
		// add custom methods here
	}

	customOmsOrderOperateHistoryModel struct {
		*defaultOmsOrderOperateHistoryModel
	}

	OmsOrderOperateHistoryRelation struct {
		//
	}
)

// NewOmsOrderOperateHistoryModel returns a model for the database table.
func NewOmsOrderOperateHistoryModel(db *gorm.DB) OmsOrderOperateHistoryModel {
	return &customOmsOrderOperateHistoryModel{
		defaultOmsOrderOperateHistoryModel: newOmsOrderOperateHistoryModel(db),
	}
}
