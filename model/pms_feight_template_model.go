package model

import (
	"gorm.io/gorm"
)

var _ PmsFeightTemplateModel = (*customPmsFeightTemplateModel)(nil)

type (
	// PmsFeightTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsFeightTemplateModel.
	PmsFeightTemplateModel interface {
		pmsFeightTemplateModel
	}

	customPmsFeightTemplateModel struct {
		*defaultPmsFeightTemplateModel
	}

	PmsFeightTemplateRelation struct {
		//
	}
)

// NewPmsFeightTemplateModel returns a model for the database table.
func NewPmsFeightTemplateModel(db *gorm.DB) PmsFeightTemplateModel {
	return &customPmsFeightTemplateModel{
		defaultPmsFeightTemplateModel: newPmsFeightTemplateModel(db),
	}
}
