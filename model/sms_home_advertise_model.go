package model

import (
	"gorm.io/gorm"
)

var _ SmsHomeAdvertiseModel = (*customSmsHomeAdvertiseModel)(nil)

type (
	// SmsHomeAdvertiseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsHomeAdvertiseModel.
	SmsHomeAdvertiseModel interface {
		smsHomeAdvertiseModel
	}

	customSmsHomeAdvertiseModel struct {
		*defaultSmsHomeAdvertiseModel
	}

	SmsHomeAdvertiseRelation struct {
		//
	}
)

// NewSmsHomeAdvertiseModel returns a model for the database table.
func NewSmsHomeAdvertiseModel(db *gorm.DB) SmsHomeAdvertiseModel {
	return &customSmsHomeAdvertiseModel{
		defaultSmsHomeAdvertiseModel: newSmsHomeAdvertiseModel(db),
	}
}
