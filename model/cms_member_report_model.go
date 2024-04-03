package model

import (
	"gorm.io/gorm"
)

var _ CmsMemberReportModel = (*customCmsMemberReportModel)(nil)

type (
	// CmsMemberReportModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsMemberReportModel.
	CmsMemberReportModel interface {
		cmsMemberReportModel
	}

	customCmsMemberReportModel struct {
		*defaultCmsMemberReportModel
	}

	CmsMemberReportRelation struct {
		//
	}
)

// NewCmsMemberReportModel returns a model for the database table.
func NewCmsMemberReportModel(db *gorm.DB) CmsMemberReportModel {
	return &customCmsMemberReportModel{
		defaultCmsMemberReportModel: newCmsMemberReportModel(db),
	}
}
