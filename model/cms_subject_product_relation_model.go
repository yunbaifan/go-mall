package model

import (
	"gorm.io/gorm"
)

var _ CmsSubjectProductRelationModel = (*customCmsSubjectProductRelationModel)(nil)

type (
	// CmsSubjectProductRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsSubjectProductRelationModel.
	//
	CmsSubjectProductRelationModel interface {
		cmsSubjectProductRelationModel
		cmsSubjectProductRelationInter
	}

	// 另外一个接口
	cmsSubjectProductRelationInter interface {
		// add custom methods here
	}

	customCmsSubjectProductRelationModel struct {
		*defaultCmsSubjectProductRelationModel
	}

	CmsSubjectProductRelationRelation struct {
		//
	}
)

// NewCmsSubjectProductRelationModel returns a model for the database table.
func NewCmsSubjectProductRelationModel(db *gorm.DB) CmsSubjectProductRelationModel {
	return &customCmsSubjectProductRelationModel{
		defaultCmsSubjectProductRelationModel: newCmsSubjectProductRelationModel(db),
	}
}
