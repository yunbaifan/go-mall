package model

import (
	"gorm.io/gorm"
)

var _ UmsGrowthChangeHistoryModel = (*customUmsGrowthChangeHistoryModel)(nil)

type (
	// UmsGrowthChangeHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsGrowthChangeHistoryModel.
	UmsGrowthChangeHistoryModel interface {
		umsGrowthChangeHistoryModel
	}

	customUmsGrowthChangeHistoryModel struct {
		*defaultUmsGrowthChangeHistoryModel
	}

	UmsGrowthChangeHistoryRelation struct {
		//
	}
)

// NewUmsGrowthChangeHistoryModel returns a model for the database table.
func NewUmsGrowthChangeHistoryModel(db *gorm.DB) UmsGrowthChangeHistoryModel {
	return &customUmsGrowthChangeHistoryModel{
		defaultUmsGrowthChangeHistoryModel: newUmsGrowthChangeHistoryModel(db),
	}
}
