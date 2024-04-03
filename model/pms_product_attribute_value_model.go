package model

import (
	"gorm.io/gorm"
)

var _ PmsProductAttributeValueModel = (*customPmsProductAttributeValueModel)(nil)

type (
	// PmsProductAttributeValueModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductAttributeValueModel.
	PmsProductAttributeValueModel interface {
		pmsProductAttributeValueModel
	}

	customPmsProductAttributeValueModel struct {
		*defaultPmsProductAttributeValueModel
	}

	PmsProductAttributeValueRelation struct {
		//
	}
)

// NewPmsProductAttributeValueModel returns a model for the database table.
func NewPmsProductAttributeValueModel(db *gorm.DB) PmsProductAttributeValueModel {
	return &customPmsProductAttributeValueModel{
		defaultPmsProductAttributeValueModel: newPmsProductAttributeValueModel(db),
	}
}
