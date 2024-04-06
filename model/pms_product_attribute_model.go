package model

import (
	"gorm.io/gorm"
)

var _ PmsProductAttributeModel = (*customPmsProductAttributeModel)(nil)

type (
	// PmsProductAttributeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductAttributeModel.
	//
	PmsProductAttributeModel interface {
		pmsProductAttributeModel
		pmsProductAttributeInter
	}

	// 另外一个接口
	pmsProductAttributeInter interface {
		// add custom methods here
	}

	customPmsProductAttributeModel struct {
		*defaultPmsProductAttributeModel
	}

	PmsProductAttributeRelation struct {
		//
	}
)

// NewPmsProductAttributeModel returns a model for the database table.
func NewPmsProductAttributeModel(db *gorm.DB) PmsProductAttributeModel {
	return &customPmsProductAttributeModel{
		defaultPmsProductAttributeModel: newPmsProductAttributeModel(db),
	}
}
