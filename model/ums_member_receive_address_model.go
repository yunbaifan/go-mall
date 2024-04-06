package model

import (
	"gorm.io/gorm"
)

var _ UmsMemberReceiveAddressModel = (*customUmsMemberReceiveAddressModel)(nil)

type (
	// UmsMemberReceiveAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberReceiveAddressModel.
	//
	UmsMemberReceiveAddressModel interface {
		umsMemberReceiveAddressModel
		umsMemberReceiveAddressInter
	}

	// 另外一个接口
	umsMemberReceiveAddressInter interface {
		// add custom methods here
	}

	customUmsMemberReceiveAddressModel struct {
		*defaultUmsMemberReceiveAddressModel
	}

	UmsMemberReceiveAddressRelation struct {
		//
	}
)

// NewUmsMemberReceiveAddressModel returns a model for the database table.
func NewUmsMemberReceiveAddressModel(db *gorm.DB) UmsMemberReceiveAddressModel {
	return &customUmsMemberReceiveAddressModel{
		defaultUmsMemberReceiveAddressModel: newUmsMemberReceiveAddressModel(db),
	}
}
