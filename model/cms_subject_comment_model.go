package model

import (
	"gorm.io/gorm"
)

var _ CmsSubjectCommentModel = (*customCmsSubjectCommentModel)(nil)

type (
	// CmsSubjectCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsSubjectCommentModel.
	//
	CmsSubjectCommentModel interface {
		cmsSubjectCommentModel
		cmsSubjectCommentInter
	}

	// 另外一个接口
	cmsSubjectCommentInter interface {
		// add custom methods here
	}

	customCmsSubjectCommentModel struct {
		*defaultCmsSubjectCommentModel
	}

	CmsSubjectCommentRelation struct {
		//
	}
)

// NewCmsSubjectCommentModel returns a model for the database table.
func NewCmsSubjectCommentModel(db *gorm.DB) CmsSubjectCommentModel {
	return &customCmsSubjectCommentModel{
		defaultCmsSubjectCommentModel: newCmsSubjectCommentModel(db),
	}
}
