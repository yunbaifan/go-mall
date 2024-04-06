package model

import (
	"gorm.io/gorm"
)

var _ PmsProductLadderModel = (*customPmsProductLadderModel)(nil)

type (
	// PmsProductLadderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductLadderModel.
	//
	PmsProductLadderModel interface {
		pmsProductLadderModel
		pmsProductLadderInter
	}

	// 另外一个接口
	pmsProductLadderInter interface {
		// add custom methods here
	}

	customPmsProductLadderModel struct {
		*defaultPmsProductLadderModel
	}

	PmsProductLadderRelation struct {
		//
	}
)

// NewPmsProductLadderModel returns a model for the database table.
func NewPmsProductLadderModel(db *gorm.DB) PmsProductLadderModel {
	return &customPmsProductLadderModel{
		defaultPmsProductLadderModel: newPmsProductLadderModel(db),
	}
}
