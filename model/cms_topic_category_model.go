package model

import (
	"gorm.io/gorm"
)

var _ CmsTopicCategoryModel = (*customCmsTopicCategoryModel)(nil)

type (
	// CmsTopicCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsTopicCategoryModel.
	//
	CmsTopicCategoryModel interface {
		cmsTopicCategoryModel
		cmsTopicCategoryInter
	}

	// 另外一个接口
	cmsTopicCategoryInter interface {
		// add custom methods here
	}

	customCmsTopicCategoryModel struct {
		*defaultCmsTopicCategoryModel
	}

	CmsTopicCategoryRelation struct {
		//
	}
)

// NewCmsTopicCategoryModel returns a model for the database table.
func NewCmsTopicCategoryModel(db *gorm.DB) CmsTopicCategoryModel {
	return &customCmsTopicCategoryModel{
		defaultCmsTopicCategoryModel: newCmsTopicCategoryModel(db),
	}
}
