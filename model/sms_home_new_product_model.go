package model

import (
	"gorm.io/gorm"
)

var _ SmsHomeNewProductModel = (*customSmsHomeNewProductModel)(nil)

type (
	// SmsHomeNewProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsHomeNewProductModel.
	SmsHomeNewProductModel interface {
		smsHomeNewProductModel
	}

	customSmsHomeNewProductModel struct {
		*defaultSmsHomeNewProductModel
	}

	SmsHomeNewProductRelation struct {
		//
	}
)

// NewSmsHomeNewProductModel returns a model for the database table.
func NewSmsHomeNewProductModel(db *gorm.DB) SmsHomeNewProductModel {
	return &customSmsHomeNewProductModel{
		defaultSmsHomeNewProductModel: newSmsHomeNewProductModel(db),
	}
}
