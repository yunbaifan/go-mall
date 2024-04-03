package model

import (
	"gorm.io/gorm"
)

var _ PmsSkuStockModel = (*customPmsSkuStockModel)(nil)

type (
	// PmsSkuStockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsSkuStockModel.
	PmsSkuStockModel interface {
		pmsSkuStockModel
	}

	customPmsSkuStockModel struct {
		*defaultPmsSkuStockModel
	}

	PmsSkuStockRelation struct {
		//
	}
)

// NewPmsSkuStockModel returns a model for the database table.
func NewPmsSkuStockModel(db *gorm.DB) PmsSkuStockModel {
	return &customPmsSkuStockModel{
		defaultPmsSkuStockModel: newPmsSkuStockModel(db),
	}
}
