package model

import (
	"gorm.io/gorm"
)

var _ UmsRoleResourceRelationModel = (*customUmsRoleResourceRelationModel)(nil)

type (
	// UmsRoleResourceRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsRoleResourceRelationModel.
	//
	UmsRoleResourceRelationModel interface {
		umsRoleResourceRelationModel
		umsRoleResourceRelationInter
	}

	// 另外一个接口
	umsRoleResourceRelationInter interface {
		// add custom methods here
	}

	customUmsRoleResourceRelationModel struct {
		*defaultUmsRoleResourceRelationModel
	}

	UmsRoleResourceRelationRelation struct {
		//
	}
)

// NewUmsRoleResourceRelationModel returns a model for the database table.
func NewUmsRoleResourceRelationModel(db *gorm.DB) UmsRoleResourceRelationModel {
	return &customUmsRoleResourceRelationModel{
		defaultUmsRoleResourceRelationModel: newUmsRoleResourceRelationModel(db),
	}
}
