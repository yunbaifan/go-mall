package model

import (
	"gorm.io/gorm"
)

var _ PmsCommentModel = (*customPmsCommentModel)(nil)

type (
	// PmsCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsCommentModel.
	//
	PmsCommentModel interface {
		pmsCommentModel
		pmsCommentInter
	}

	// 另外一个接口
	pmsCommentInter interface {
		// add custom methods here
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
