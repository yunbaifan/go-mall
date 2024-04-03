package model

import (
	"gorm.io/gorm"
)

var _ OmsCompanyAddressModel = (*customOmsCompanyAddressModel)(nil)

type (
	// OmsCompanyAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsCompanyAddressModel.
	OmsCompanyAddressModel interface {
		omsCompanyAddressModel
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
