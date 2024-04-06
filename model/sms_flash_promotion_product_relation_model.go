package model

import (
	"gorm.io/gorm"
)

var _ SmsFlashPromotionProductRelationModel = (*customSmsFlashPromotionProductRelationModel)(nil)

type (
	// SmsFlashPromotionProductRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsFlashPromotionProductRelationModel.
	//
	SmsFlashPromotionProductRelationModel interface {
		smsFlashPromotionProductRelationModel
		smsFlashPromotionProductRelationInter
	}

	// 另外一个接口
	smsFlashPromotionProductRelationInter interface {
		// add custom methods here
	}

	customSmsFlashPromotionProductRelationModel struct {
		*defaultSmsFlashPromotionProductRelationModel
	}

	SmsFlashPromotionProductRelationRelation struct {
		//
	}
)

// NewSmsFlashPromotionProductRelationModel returns a model for the database table.
func NewSmsFlashPromotionProductRelationModel(db *gorm.DB) SmsFlashPromotionProductRelationModel {
	return &customSmsFlashPromotionProductRelationModel{
		defaultSmsFlashPromotionProductRelationModel: newSmsFlashPromotionProductRelationModel(db),
	}
}
