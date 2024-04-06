package model

import (
	"gorm.io/gorm"
)

var _ SmsHomeAdvertiseModel = (*customSmsHomeAdvertiseModel)(nil)

type (
	// SmsHomeAdvertiseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsHomeAdvertiseModel.
	//
	SmsHomeAdvertiseModel interface {
		smsHomeAdvertiseModel
		smsHomeAdvertiseInter
	}

	// 另外一个接口
	smsHomeAdvertiseInter interface {
		// add custom methods here
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
