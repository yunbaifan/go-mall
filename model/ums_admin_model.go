package model

import (
	"context"
	"gorm.io/gorm"
)

var _ UmsAdminModel = (*customUmsAdminModel)(nil)

type (
	// UmsAdminModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsAdminModel.
	UmsAdminModel interface {
		umsAdminModel
		umsAdminInter
	}

	umsAdminInter interface {
		FindOneByUserName(ctx context.Context, username string) (*UmsAdmin, error)
	}

	customUmsAdminModel struct {
		*defaultUmsAdminModel
	}

	UmsAdminRelation struct {
		//
	}
)

// NewUmsAdminModel returns a model for the database table.
func NewUmsAdminModel(db *gorm.DB) UmsAdminModel {
	return &customUmsAdminModel{
		defaultUmsAdminModel: newUmsAdminModel(db),
	}
}

func (d *defaultUmsAdminModel) FindOneByUserName(ctx context.Context, username string) (*UmsAdmin, error) {
	var mm UmsAdmin
	err := d.OrmSession(ctx).Where("username = ?", username).First(&mm).Error
	switch err {
	case nil:
		return &mm, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}
