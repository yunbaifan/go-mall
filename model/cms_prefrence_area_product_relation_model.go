package model

import (
	"gorm.io/gorm"
)

var _ CmsPrefrenceAreaProductRelationModel = (*customCmsPrefrenceAreaProductRelationModel)(nil)

type (
	// CmsPrefrenceAreaProductRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsPrefrenceAreaProductRelationModel.
	//
	CmsPrefrenceAreaProductRelationModel interface {
		cmsPrefrenceAreaProductRelationModel
		cmsPrefrenceAreaProductRelationInter
	}

	// 另外一个接口
	cmsPrefrenceAreaProductRelationInter interface {
		// add custom methods here
	}

	customCmsPrefrenceAreaProductRelationModel struct {
		*defaultCmsPrefrenceAreaProductRelationModel
	}

	CmsPrefrenceAreaProductRelationRelation struct {
		//
	}
)

// NewCmsPrefrenceAreaProductRelationModel returns a model for the database table.
func NewCmsPrefrenceAreaProductRelationModel(db *gorm.DB) CmsPrefrenceAreaProductRelationModel {
	return &customCmsPrefrenceAreaProductRelationModel{
		defaultCmsPrefrenceAreaProductRelationModel: newCmsPrefrenceAreaProductRelationModel(db),
	}
}
