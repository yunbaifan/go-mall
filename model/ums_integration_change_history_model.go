package model

import (
	"gorm.io/gorm"
)

var _ UmsIntegrationChangeHistoryModel = (*customUmsIntegrationChangeHistoryModel)(nil)

type (
	// UmsIntegrationChangeHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsIntegrationChangeHistoryModel.
	UmsIntegrationChangeHistoryModel interface {
		umsIntegrationChangeHistoryModel
	}

	customUmsIntegrationChangeHistoryModel struct {
		*defaultUmsIntegrationChangeHistoryModel
	}

	UmsIntegrationChangeHistoryRelation struct {
		//
	}
)

// NewUmsIntegrationChangeHistoryModel returns a model for the database table.
func NewUmsIntegrationChangeHistoryModel(db *gorm.DB) UmsIntegrationChangeHistoryModel {
	return &customUmsIntegrationChangeHistoryModel{
		defaultUmsIntegrationChangeHistoryModel: newUmsIntegrationChangeHistoryModel(db),
	}
}
