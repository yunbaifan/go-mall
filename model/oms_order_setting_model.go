package model

import (
	"gorm.io/gorm"
)

var _ OmsOrderSettingModel = (*customOmsOrderSettingModel)(nil)

type (
	// OmsOrderSettingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderSettingModel.
	//
	OmsOrderSettingModel interface {
		omsOrderSettingModel
		omsOrderSettingInter
	}

	// 另外一个接口
	omsOrderSettingInter interface {
		// add custom methods here
	}

	customOmsOrderSettingModel struct {
		*defaultOmsOrderSettingModel
	}

	OmsOrderSettingRelation struct {
		//
	}
)

// NewOmsOrderSettingModel returns a model for the database table.
func NewOmsOrderSettingModel(db *gorm.DB) OmsOrderSettingModel {
	return &customOmsOrderSettingModel{
		defaultOmsOrderSettingModel: newOmsOrderSettingModel(db),
	}
}
