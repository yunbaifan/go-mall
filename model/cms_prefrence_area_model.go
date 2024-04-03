package model

import (
	"gorm.io/gorm"
)

var _ CmsPrefrenceAreaModel = (*customCmsPrefrenceAreaModel)(nil)

type (
	// CmsPrefrenceAreaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsPrefrenceAreaModel.
	CmsPrefrenceAreaModel interface {
		cmsPrefrenceAreaModel
	}

	customCmsPrefrenceAreaModel struct {
		*defaultCmsPrefrenceAreaModel
	}

	CmsPrefrenceAreaRelation struct {
		//
	}
)

// NewCmsPrefrenceAreaModel returns a model for the database table.
func NewCmsPrefrenceAreaModel(db *gorm.DB) CmsPrefrenceAreaModel {
	return &customCmsPrefrenceAreaModel{
		defaultCmsPrefrenceAreaModel: newCmsPrefrenceAreaModel(db),
	}
}
