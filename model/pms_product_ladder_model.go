package model

import (
	"gorm.io/gorm"
)

var _ PmsProductLadderModel = (*customPmsProductLadderModel)(nil)

type (
	// PmsProductLadderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductLadderModel.
	PmsProductLadderModel interface {
		pmsProductLadderModel
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
