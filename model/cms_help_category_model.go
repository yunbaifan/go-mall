package model

import (
	"gorm.io/gorm"
)

var _ CmsHelpCategoryModel = (*customCmsHelpCategoryModel)(nil)

type (
	// CmsHelpCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsHelpCategoryModel.
	//
	CmsHelpCategoryModel interface {
		cmsHelpCategoryModel
		cmsHelpCategoryInter
	}

	// 另外一个接口
	cmsHelpCategoryInter interface {
		// add custom methods here
	}

	customCmsHelpCategoryModel struct {
		*defaultCmsHelpCategoryModel
	}

	CmsHelpCategoryRelation struct {
		//
	}
)

// NewCmsHelpCategoryModel returns a model for the database table.
func NewCmsHelpCategoryModel(db *gorm.DB) CmsHelpCategoryModel {
	return &customCmsHelpCategoryModel{
		defaultCmsHelpCategoryModel: newCmsHelpCategoryModel(db),
	}
}
