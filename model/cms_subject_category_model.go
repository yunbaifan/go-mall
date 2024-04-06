package model

import (
	"gorm.io/gorm"
)

var _ CmsSubjectCategoryModel = (*customCmsSubjectCategoryModel)(nil)

type (
	// CmsSubjectCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsSubjectCategoryModel.
	//
	CmsSubjectCategoryModel interface {
		cmsSubjectCategoryModel
		cmsSubjectCategoryInter
	}

	// 另外一个接口
	cmsSubjectCategoryInter interface {
		// add custom methods here
	}

	customCmsSubjectCategoryModel struct {
		*defaultCmsSubjectCategoryModel
	}

	CmsSubjectCategoryRelation struct {
		//
	}
)

// NewCmsSubjectCategoryModel returns a model for the database table.
func NewCmsSubjectCategoryModel(db *gorm.DB) CmsSubjectCategoryModel {
	return &customCmsSubjectCategoryModel{
		defaultCmsSubjectCategoryModel: newCmsSubjectCategoryModel(db),
	}
}
