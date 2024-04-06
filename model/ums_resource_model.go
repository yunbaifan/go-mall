package model

import (
	"gorm.io/gorm"
)

var _ UmsResourceModel = (*customUmsResourceModel)(nil)

type (
	// UmsResourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsResourceModel.
	//
	UmsResourceModel interface {
		umsResourceModel
		umsResourceInter
	}

	// 另外一个接口
	umsResourceInter interface {
		// add custom methods here
	}

	customUmsResourceModel struct {
		*defaultUmsResourceModel
	}

	UmsResourceRelation struct {
		//
	}
)

// NewUmsResourceModel returns a model for the database table.
func NewUmsResourceModel(db *gorm.DB) UmsResourceModel {
	return &customUmsResourceModel{
		defaultUmsResourceModel: newUmsResourceModel(db),
	}
}
