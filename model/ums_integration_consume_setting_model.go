package model

import (
	"gorm.io/gorm"
)

var _ UmsIntegrationConsumeSettingModel = (*customUmsIntegrationConsumeSettingModel)(nil)

type (
	// UmsIntegrationConsumeSettingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsIntegrationConsumeSettingModel.
	UmsIntegrationConsumeSettingModel interface {
		umsIntegrationConsumeSettingModel
	}

	customUmsIntegrationConsumeSettingModel struct {
		*defaultUmsIntegrationConsumeSettingModel
	}

	UmsIntegrationConsumeSettingRelation struct {
		//
	}
)

// NewUmsIntegrationConsumeSettingModel returns a model for the database table.
func NewUmsIntegrationConsumeSettingModel(db *gorm.DB) UmsIntegrationConsumeSettingModel {
	return &customUmsIntegrationConsumeSettingModel{
		defaultUmsIntegrationConsumeSettingModel: newUmsIntegrationConsumeSettingModel(db),
	}
}
