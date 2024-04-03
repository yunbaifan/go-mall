package model

import (
	"gorm.io/gorm"
)

var _ CmsTopicCommentModel = (*customCmsTopicCommentModel)(nil)

type (
	// CmsTopicCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsTopicCommentModel.
	CmsTopicCommentModel interface {
		cmsTopicCommentModel
	}

	customCmsTopicCommentModel struct {
		*defaultCmsTopicCommentModel
	}

	CmsTopicCommentRelation struct {
		//
	}
)

// NewCmsTopicCommentModel returns a model for the database table.
func NewCmsTopicCommentModel(db *gorm.DB) CmsTopicCommentModel {
	return &customCmsTopicCommentModel{
		defaultCmsTopicCommentModel: newCmsTopicCommentModel(db),
	}
}
