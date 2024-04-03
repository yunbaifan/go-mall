package model

import (
	"gorm.io/gorm"
)

var _ UmsMemberRuleSettingModel = (*customUmsMemberRuleSettingModel)(nil)

type (
	// UmsMemberRuleSettingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberRuleSettingModel.
	UmsMemberRuleSettingModel interface {
		umsMemberRuleSettingModel
	}

	customUmsMemberRuleSettingModel struct {
		*defaultUmsMemberRuleSettingModel
	}

	UmsMemberRuleSettingRelation struct {
		//
	}
)

// NewUmsMemberRuleSettingModel returns a model for the database table.
func NewUmsMemberRuleSettingModel(db *gorm.DB) UmsMemberRuleSettingModel {
	return &customUmsMemberRuleSettingModel{
		defaultUmsMemberRuleSettingModel: newUmsMemberRuleSettingModel(db),
	}
}
