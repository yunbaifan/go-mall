package model

import (
	"gorm.io/gorm"
)

var _ PmsProductFullReductionModel = (*customPmsProductFullReductionModel)(nil)

type (
	// PmsProductFullReductionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductFullReductionModel.
	//
	PmsProductFullReductionModel interface {
		pmsProductFullReductionModel
		pmsProductFullReductionInter
	}

	// 另外一个接口
	pmsProductFullReductionInter interface {
		// add custom methods here
	}

	customPmsProductFullReductionModel struct {
		*defaultPmsProductFullReductionModel
	}

	PmsProductFullReductionRelation struct {
		//
	}
)

// NewPmsProductFullReductionModel returns a model for the database table.
func NewPmsProductFullReductionModel(db *gorm.DB) PmsProductFullReductionModel {
	return &customPmsProductFullReductionModel{
		defaultPmsProductFullReductionModel: newPmsProductFullReductionModel(db),
	}
}
