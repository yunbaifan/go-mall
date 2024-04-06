package model

import (
	"gorm.io/gorm"
)

var _ OmsCompanyAddressModel = (*customOmsCompanyAddressModel)(nil)

type (
	// OmsCompanyAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsCompanyAddressModel.
	//
	OmsCompanyAddressModel interface {
		omsCompanyAddressModel
		omsCompanyAddressInter
	}

	// 另外一个接口
	omsCompanyAddressInter interface {
		// add custom methods here
	}

	customOmsCompanyAddressModel struct {
		*defaultOmsCompanyAddressModel
	}

	OmsCompanyAddressRelation struct {
		//
	}
)

// NewOmsCompanyAddressModel returns a model for the database table.
func NewOmsCompanyAddressModel(db *gorm.DB) OmsCompanyAddressModel {
	return &customOmsCompanyAddressModel{
		defaultOmsCompanyAddressModel: newOmsCompanyAddressModel(db),
	}
}
