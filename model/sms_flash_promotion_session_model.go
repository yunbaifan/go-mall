package model

import (
	"gorm.io/gorm"
)

var _ SmsFlashPromotionSessionModel = (*customSmsFlashPromotionSessionModel)(nil)

type (
	// SmsFlashPromotionSessionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsFlashPromotionSessionModel.
	SmsFlashPromotionSessionModel interface {
		smsFlashPromotionSessionModel
	}

	customSmsFlashPromotionSessionModel struct {
		*defaultSmsFlashPromotionSessionModel
	}

	SmsFlashPromotionSessionRelation struct {
		//
	}
)

// NewSmsFlashPromotionSessionModel returns a model for the database table.
func NewSmsFlashPromotionSessionModel(db *gorm.DB) SmsFlashPromotionSessionModel {
	return &customSmsFlashPromotionSessionModel{
		defaultSmsFlashPromotionSessionModel: newSmsFlashPromotionSessionModel(db),
	}
}
