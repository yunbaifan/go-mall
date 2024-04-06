package model

import (
	"gorm.io/gorm"
)

var _ SmsHomeRecommendProductModel = (*customSmsHomeRecommendProductModel)(nil)

type (
	// SmsHomeRecommendProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsHomeRecommendProductModel.
	//
	SmsHomeRecommendProductModel interface {
		smsHomeRecommendProductModel
		smsHomeRecommendProductInter
	}

	// 另外一个接口
	smsHomeRecommendProductInter interface {
		// add custom methods here
	}

	customSmsHomeRecommendProductModel struct {
		*defaultSmsHomeRecommendProductModel
	}

	SmsHomeRecommendProductRelation struct {
		//
	}
)

// NewSmsHomeRecommendProductModel returns a model for the database table.
func NewSmsHomeRecommendProductModel(db *gorm.DB) SmsHomeRecommendProductModel {
	return &customSmsHomeRecommendProductModel{
		defaultSmsHomeRecommendProductModel: newSmsHomeRecommendProductModel(db),
	}
}
