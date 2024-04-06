package model

import (
	"gorm.io/gorm"
)

var _ PmsAlbumModel = (*customPmsAlbumModel)(nil)

type (
	// PmsAlbumModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsAlbumModel.
	//
	PmsAlbumModel interface {
		pmsAlbumModel
		pmsAlbumInter
	}

	// 另外一个接口
	pmsAlbumInter interface {
		// add custom methods here
	}

	customPmsAlbumModel struct {
		*defaultPmsAlbumModel
	}

	PmsAlbumRelation struct {
		//
	}
)

// NewPmsAlbumModel returns a model for the database table.
func NewPmsAlbumModel(db *gorm.DB) PmsAlbumModel {
	return &customPmsAlbumModel{
		defaultPmsAlbumModel: newPmsAlbumModel(db),
	}
}
