package model

import (
	"gorm.io/gorm"
)

var _ CmsPrefrenceAreaProductRelationModel = (*customCmsPrefrenceAreaProductRelationModel)(nil)

type (
	// CmsPrefrenceAreaProductRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsPrefrenceAreaProductRelationModel.
	CmsPrefrenceAreaProductRelationModel interface {
		cmsPrefrenceAreaProductRelationModel
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
