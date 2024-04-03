package model

import (
	"gorm.io/gorm"
)

var _ PmsProductOperateLogModel = (*customPmsProductOperateLogModel)(nil)

type (
	// PmsProductOperateLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductOperateLogModel.
	PmsProductOperateLogModel interface {
		pmsProductOperateLogModel
	}

	customPmsProductOperateLogModel struct {
		*defaultPmsProductOperateLogModel
	}

	PmsProductOperateLogRelation struct {
		//
	}
)

// NewPmsProductOperateLogModel returns a model for the database table.
func NewPmsProductOperateLogModel(db *gorm.DB) PmsProductOperateLogModel {
	return &customPmsProductOperateLogModel{
		defaultPmsProductOperateLogModel: newPmsProductOperateLogModel(db),
	}
}
