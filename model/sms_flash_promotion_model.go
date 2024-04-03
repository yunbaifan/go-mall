package model

import (
	"gorm.io/gorm"
)

var _ SmsFlashPromotionModel = (*customSmsFlashPromotionModel)(nil)

type (
	// SmsFlashPromotionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsFlashPromotionModel.
	SmsFlashPromotionModel interface {
		smsFlashPromotionModel
	}

	customSmsFlashPromotionModel struct {
		*defaultSmsFlashPromotionModel
	}

	SmsFlashPromotionRelation struct {
		//
	}
)

// NewSmsFlashPromotionModel returns a model for the database table.
func NewSmsFlashPromotionModel(db *gorm.DB) SmsFlashPromotionModel {
	return &customSmsFlashPromotionModel{
		defaultSmsFlashPromotionModel: newSmsFlashPromotionModel(db),
	}
}
