package model

import (
	"gorm.io/gorm"
)

var _ PmsProductVertifyRecordModel = (*customPmsProductVertifyRecordModel)(nil)

type (
	// PmsProductVertifyRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductVertifyRecordModel.
	PmsProductVertifyRecordModel interface {
		pmsProductVertifyRecordModel
	}

	customPmsProductVertifyRecordModel struct {
		*defaultPmsProductVertifyRecordModel
	}

	PmsProductVertifyRecordRelation struct {
		//
	}
)

// NewPmsProductVertifyRecordModel returns a model for the database table.
func NewPmsProductVertifyRecordModel(db *gorm.DB) PmsProductVertifyRecordModel {
	return &customPmsProductVertifyRecordModel{
		defaultPmsProductVertifyRecordModel: newPmsProductVertifyRecordModel(db),
	}
}
