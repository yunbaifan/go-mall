// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"math"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var (
	_ sql.IsolationLevel
)

type (
	smsHomeBrandModel interface {
		Insert(ctx context.Context, data *SmsHomeBrand) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*SmsHomeBrand, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, page, pageSize int64) (*SmsHomeBrandPagination, error)
		Update(ctx context.Context, data *SmsHomeBrand) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSmsHomeBrandModel struct {
		table   string
		ormConn *gorm.DB
	}

	SmsHomeBrand struct {
		Id              int64          `json:"id" gorm:"primary_key"`
		BrandId         sql.NullInt64  `json:"brand_id"`
		BrandName       sql.NullString `json:"brand_name"`
		RecommendStatus sql.NullInt64  `json:"recommend_status"`
		Sort            sql.NullInt64  `json:"sort"`
		SmsHomeBrandRelation
	}

	SmsHomeBrandPagination struct {
		Data []*SmsHomeBrand
		SmsHomeBrandPaginationBasic
	}

	SmsHomeBrandPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}
)

func newSmsHomeBrandModel(db *gorm.DB) *defaultSmsHomeBrandModel {
	return &defaultSmsHomeBrandModel{
		table:   "`sms_home_brand`",
		ormConn: db,
	}
}

func (SmsHomeBrand) TableName() string {
	return "`sms_home_brand`"
}

func (m SmsHomeBrand) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultSmsHomeBrandModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&SmsHomeBrand{}, id).Error
}

func (m *defaultSmsHomeBrandModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*SmsHomeBrand, error) {
	var mm SmsHomeBrand
	db := m.OrmSession(ctx)
	for _, preload := range preloadList {
		db = db.Preload(preload)
	}
	err := db.First(&mm, id).Error
	switch err {
	case nil:
		return &mm, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

func (m *defaultSmsHomeBrandModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&SmsHomeBrand{}).WithContext(ctx)
}

func (m *defaultSmsHomeBrandModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultSmsHomeBrandModel) FindPageListByBuilder(db *gorm.DB, page, pageSize int64) (*SmsHomeBrandPagination, error) {
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(SmsHomeBrandPagination)
	resp.PageSize = pageSize
	resp.TotalPage = totalPage
	resp.CurrentPage = page
	resp.TotalCount = totalCount

	// 没有数据则返回空数据
	if totalCount <= 0 {
		return resp, nil
	}
	// 当前页数没有数据也返回空数据
	if totalCount < ((page - 1) * pageSize) {
		return resp, nil
	}

	offset := int((page - 1) * pageSize)
	if err := db.Offset(offset).Limit(int(pageSize)).Find(&resp.Data).Error; err != nil {
		return nil, err
	}

	return resp, nil
}

func (m SmsHomeBrandPagination) Pagination() SmsHomeBrandPaginationBasic {
	return SmsHomeBrandPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m SmsHomeBrandPagination) ToProto(to any, toData any, toPagination any) error {
	if to != nil {
		if err := copier.CopyWithOption(to, m, copier.Option{}); err != nil {
			return err
		}
	}
	if toData != nil {
		if err := copier.Copy(toData, m.Data); err != nil {
			return err
		}
	}
	if toPagination != nil {
		if err := copier.Copy(toPagination, m.Pagination()); err != nil {
			return err
		}
	}
	return nil
}

func (m *defaultSmsHomeBrandModel) Insert(ctx context.Context, data *SmsHomeBrand) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultSmsHomeBrandModel) Update(ctx context.Context, data *SmsHomeBrand) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
