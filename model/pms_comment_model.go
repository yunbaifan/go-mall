package model

import (
	"gorm.io/gorm"
)

var _ PmsCommentModel = (*customPmsCommentModel)(nil)

type (
	// PmsCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsCommentModel.
	PmsCommentModel interface {
		pmsCommentModel
	}

	customPmsCommentModel struct {
		*defaultPmsCommentModel
	}

	PmsCommentRelation struct {
		//
	}
)

// NewPmsCommentModel returns a model for the database table.
func NewPmsCommentModel(db *gorm.DB) PmsCommentModel {
	return &customPmsCommentModel{
		defaultPmsCommentModel: newPmsCommentModel(db),
	}
}
