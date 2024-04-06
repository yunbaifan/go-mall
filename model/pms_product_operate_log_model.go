package model

import (
	"gorm.io/gorm"
)

var _ PmsProductOperateLogModel = (*customPmsProductOperateLogModel)(nil)

type (
	// PmsProductOperateLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductOperateLogModel.
	//
	PmsProductOperateLogModel interface {
		pmsProductOperateLogModel
		pmsProductOperateLogInter
	}

	// 另外一个接口
	pmsProductOperateLogInter interface {
		// add custom methods here
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
