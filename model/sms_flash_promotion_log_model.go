package model

import (
	"gorm.io/gorm"
)

var _ SmsFlashPromotionLogModel = (*customSmsFlashPromotionLogModel)(nil)

type (
	// SmsFlashPromotionLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsFlashPromotionLogModel.
	//
	SmsFlashPromotionLogModel interface {
		smsFlashPromotionLogModel
		smsFlashPromotionLogInter
	}

	// 另外一个接口
	smsFlashPromotionLogInter interface {
		// add custom methods here
	}

	customSmsFlashPromotionLogModel struct {
		*defaultSmsFlashPromotionLogModel
	}

	SmsFlashPromotionLogRelation struct {
		//
	}
)

// NewSmsFlashPromotionLogModel returns a model for the database table.
func NewSmsFlashPromotionLogModel(db *gorm.DB) SmsFlashPromotionLogModel {
	return &customSmsFlashPromotionLogModel{
		defaultSmsFlashPromotionLogModel: newSmsFlashPromotionLogModel(db),
	}
}
