package model

import (
	"gorm.io/gorm"
)

var _ OmsCartItemModel = (*customOmsCartItemModel)(nil)

type (
	// OmsCartItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsCartItemModel.
	//
	OmsCartItemModel interface {
		omsCartItemModel
		omsCartItemInter
	}

	// 另外一个接口
	omsCartItemInter interface {
		// add custom methods here
	}

	customOmsCartItemModel struct {
		*defaultOmsCartItemModel
	}

	OmsCartItemRelation struct {
		//
	}
)

// NewOmsCartItemModel returns a model for the database table.
func NewOmsCartItemModel(db *gorm.DB) OmsCartItemModel {
	return &customOmsCartItemModel{
		defaultOmsCartItemModel: newOmsCartItemModel(db),
	}
}
