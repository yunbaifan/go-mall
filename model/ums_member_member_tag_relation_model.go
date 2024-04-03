package model

import (
	"gorm.io/gorm"
)

var _ UmsMemberMemberTagRelationModel = (*customUmsMemberMemberTagRelationModel)(nil)

type (
	// UmsMemberMemberTagRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberMemberTagRelationModel.
	UmsMemberMemberTagRelationModel interface {
		umsMemberMemberTagRelationModel
	}

	customUmsMemberMemberTagRelationModel struct {
		*defaultUmsMemberMemberTagRelationModel
	}

	UmsMemberMemberTagRelationRelation struct {
		//
	}
)

// NewUmsMemberMemberTagRelationModel returns a model for the database table.
func NewUmsMemberMemberTagRelationModel(db *gorm.DB) UmsMemberMemberTagRelationModel {
	return &customUmsMemberMemberTagRelationModel{
		defaultUmsMemberMemberTagRelationModel: newUmsMemberMemberTagRelationModel(db),
	}
}
