package model

import (
	"gorm.io/gorm"
)

var _ CmsTopicModel = (*customCmsTopicModel)(nil)

type (
	// CmsTopicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsTopicModel.
	//
	CmsTopicModel interface {
		cmsTopicModel
		cmsTopicInter
	}

	// 另外一个接口
	cmsTopicInter interface {
		// add custom methods here
	}

	customCmsTopicModel struct {
		*defaultCmsTopicModel
	}

	CmsTopicRelation struct {
		//
	}
)

// NewCmsTopicModel returns a model for the database table.
func NewCmsTopicModel(db *gorm.DB) CmsTopicModel {
	return &customCmsTopicModel{
		defaultCmsTopicModel: newCmsTopicModel(db),
	}
}
