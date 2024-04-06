package model

import (
	"gorm.io/gorm"
)

var _ PmsAlbumPicModel = (*customPmsAlbumPicModel)(nil)

type (
	// PmsAlbumPicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsAlbumPicModel.
	//
	PmsAlbumPicModel interface {
		pmsAlbumPicModel
		pmsAlbumPicInter
	}

	// 另外一个接口
	pmsAlbumPicInter interface {
		// add custom methods here
	}

	customPmsAlbumPicModel struct {
		*defaultPmsAlbumPicModel
	}

	PmsAlbumPicRelation struct {
		//
	}
)

// NewPmsAlbumPicModel returns a model for the database table.
func NewPmsAlbumPicModel(db *gorm.DB) PmsAlbumPicModel {
	return &customPmsAlbumPicModel{
		defaultPmsAlbumPicModel: newPmsAlbumPicModel(db),
	}
}
