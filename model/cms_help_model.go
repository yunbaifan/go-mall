package model

import (
	"gorm.io/gorm"
)

var _ CmsHelpModel = (*customCmsHelpModel)(nil)

type (
	// CmsHelpModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsHelpModel.
	//
	CmsHelpModel interface {
		cmsHelpModel
		cmsHelpInter
	}

	// 另外一个接口
	cmsHelpInter interface {
		// add custom methods here
	}

	customCmsHelpModel struct {
		*defaultCmsHelpModel
	}

	CmsHelpRelation struct {
		//
	}
)

// NewCmsHelpModel returns a model for the database table.
func NewCmsHelpModel(db *gorm.DB) CmsHelpModel {
	return &customCmsHelpModel{
		defaultCmsHelpModel: newCmsHelpModel(db),
	}
}
