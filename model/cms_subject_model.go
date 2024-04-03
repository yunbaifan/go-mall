package model

import (
	"gorm.io/gorm"
)

var _ CmsSubjectModel = (*customCmsSubjectModel)(nil)

type (
	// CmsSubjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsSubjectModel.
	CmsSubjectModel interface {
		cmsSubjectModel
	}

	customCmsSubjectModel struct {
		*defaultCmsSubjectModel
	}

	CmsSubjectRelation struct {
		//
	}
)

// NewCmsSubjectModel returns a model for the database table.
func NewCmsSubjectModel(db *gorm.DB) CmsSubjectModel {
	return &customCmsSubjectModel{
		defaultCmsSubjectModel: newCmsSubjectModel(db),
	}
}
