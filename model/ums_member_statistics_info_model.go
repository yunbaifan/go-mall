package model

import (
	"gorm.io/gorm"
)

var _ UmsMemberStatisticsInfoModel = (*customUmsMemberStatisticsInfoModel)(nil)

type (
	// UmsMemberStatisticsInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberStatisticsInfoModel.
	//
	UmsMemberStatisticsInfoModel interface {
		umsMemberStatisticsInfoModel
		umsMemberStatisticsInfoInter
	}

	// 另外一个接口
	umsMemberStatisticsInfoInter interface {
		// add custom methods here
	}

	customUmsMemberStatisticsInfoModel struct {
		*defaultUmsMemberStatisticsInfoModel
	}

	UmsMemberStatisticsInfoRelation struct {
		//
	}
)

// NewUmsMemberStatisticsInfoModel returns a model for the database table.
func NewUmsMemberStatisticsInfoModel(db *gorm.DB) UmsMemberStatisticsInfoModel {
	return &customUmsMemberStatisticsInfoModel{
		defaultUmsMemberStatisticsInfoModel: newUmsMemberStatisticsInfoModel(db),
	}
}
