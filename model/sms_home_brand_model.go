package model

import (
	"gorm.io/gorm"
)

var _ SmsHomeBrandModel = (*customSmsHomeBrandModel)(nil)

type (
	// SmsHomeBrandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsHomeBrandModel.
	//
	SmsHomeBrandModel interface {
		smsHomeBrandModel
		smsHomeBrandInter
	}

	// 另外一个接口
	smsHomeBrandInter interface {
		// add custom methods here
	}

	customSmsHomeBrandModel struct {
		*defaultSmsHomeBrandModel
	}

	SmsHomeBrandRelation struct {
		//
	}
)

// NewSmsHomeBrandModel returns a model for the database table.
func NewSmsHomeBrandModel(db *gorm.DB) SmsHomeBrandModel {
	return &customSmsHomeBrandModel{
		defaultSmsHomeBrandModel: newSmsHomeBrandModel(db),
	}
}
