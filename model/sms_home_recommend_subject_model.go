package model

import (
	"gorm.io/gorm"
)

var _ SmsHomeRecommendSubjectModel = (*customSmsHomeRecommendSubjectModel)(nil)

type (
	// SmsHomeRecommendSubjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsHomeRecommendSubjectModel.
	SmsHomeRecommendSubjectModel interface {
		smsHomeRecommendSubjectModel
	}

	customSmsHomeRecommendSubjectModel struct {
		*defaultSmsHomeRecommendSubjectModel
	}

	SmsHomeRecommendSubjectRelation struct {
		//
	}
)

// NewSmsHomeRecommendSubjectModel returns a model for the database table.
func NewSmsHomeRecommendSubjectModel(db *gorm.DB) SmsHomeRecommendSubjectModel {
	return &customSmsHomeRecommendSubjectModel{
		defaultSmsHomeRecommendSubjectModel: newSmsHomeRecommendSubjectModel(db),
	}
}
