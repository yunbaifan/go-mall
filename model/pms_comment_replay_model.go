package model

import (
	"gorm.io/gorm"
)

var _ PmsCommentReplayModel = (*customPmsCommentReplayModel)(nil)

type (
	// PmsCommentReplayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsCommentReplayModel.
	//
	PmsCommentReplayModel interface {
		pmsCommentReplayModel
		pmsCommentReplayInter
	}

	// 另外一个接口
	pmsCommentReplayInter interface {
		// add custom methods here
	}

	customPmsCommentReplayModel struct {
		*defaultPmsCommentReplayModel
	}

	PmsCommentReplayRelation struct {
		//
	}
)

// NewPmsCommentReplayModel returns a model for the database table.
func NewPmsCommentReplayModel(db *gorm.DB) PmsCommentReplayModel {
	return &customPmsCommentReplayModel{
		defaultPmsCommentReplayModel: newPmsCommentReplayModel(db),
	}
}
