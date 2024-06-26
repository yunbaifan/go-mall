package model

import (
	"gorm.io/gorm"
)

var _ UmsRoleModel = (*customUmsRoleModel)(nil)

type (
	// UmsRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsRoleModel.
	//
	UmsRoleModel interface {
		umsRoleModel
		umsRoleInter
	}

	// 另外一个接口
	umsRoleInter interface {
		// add custom methods here
	}

	customUmsRoleModel struct {
		*defaultUmsRoleModel
	}

	UmsRoleRelation struct {
		//
	}
)

// NewUmsRoleModel returns a model for the database table.
func NewUmsRoleModel(db *gorm.DB) UmsRoleModel {
	return &customUmsRoleModel{
		defaultUmsRoleModel: newUmsRoleModel(db),
	}
}
