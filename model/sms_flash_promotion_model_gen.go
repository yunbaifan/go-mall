// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"math"

	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var (
	_ sql.IsolationLevel
)

type (
	smsFlashPromotionModel interface {
		Insert(ctx context.Context, data *SmsFlashPromotion) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*SmsFlashPromotion, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ctx context.Context, keyword *KeywordSmsFlashPromotionModel) (*SmsFlashPromotionPagination, error)
		Update(ctx context.Context, data *SmsFlashPromotion) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSmsFlashPromotionModel struct {
		table   string
		ormConn *gorm.DB
	}

	SmsFlashPromotion struct {
		Id         int64          `json:"id" gorm:"primary_key"`
		Title      sql.NullString `json:"title"`       // 秒杀时间段名称
		StartDate  sql.NullTime   `json:"start_date"`  // 开始日期
		EndDate    sql.NullTime   `json:"end_date"`    // 结束日期
		Status     sql.NullInt64  `json:"status"`      // 上下线状态
		CreateTime sql.NullTime   `json:"create_time"` // 创建时间
		SmsFlashPromotionRelation
	}

	SmsFlashPromotionPagination struct {
		Data []*SmsFlashPromotion
		SmsFlashPromotionPaginationBasic
	}

	SmsFlashPromotionPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordSmsFlashPromotionModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newSmsFlashPromotionModel(db *gorm.DB) *defaultSmsFlashPromotionModel {
	return &defaultSmsFlashPromotionModel{
		table:   "`sms_flash_promotion`",
		ormConn: db,
	}
}

func (SmsFlashPromotion) TableName() string {
	return "`sms_flash_promotion`"
}

func (m SmsFlashPromotion) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultSmsFlashPromotionModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&SmsFlashPromotion{}, id).Error
}

func (m *defaultSmsFlashPromotionModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*SmsFlashPromotion, error) {
	var mm SmsFlashPromotion
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

func (m *defaultSmsFlashPromotionModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&SmsFlashPromotion{}).WithContext(ctx)
}

func (m *defaultSmsFlashPromotionModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultSmsFlashPromotionModel) FindPageListByBuilder(ctx context.Context, keyword *KeywordSmsFlashPromotionModel) (*SmsFlashPromotionPagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	db := m.OrmSession(ctx)
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(SmsFlashPromotionPagination)
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
	if keyword.KeywordKey != "" && keyword.KeywordValue != "" {
		db = db.Where(fmt.Sprintf("%s = ?", keyword.KeywordKey), keyword.KeywordValue)
	}
	offset := int((page - 1) * pageSize)
	if err := db.Offset(offset).Limit(int(pageSize)).Find(&resp.Data).Error; err != nil {
		return nil, err
	}

	return resp, nil
}

func (m SmsFlashPromotionPagination) Pagination() SmsFlashPromotionPaginationBasic {
	return SmsFlashPromotionPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m SmsFlashPromotionPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultSmsFlashPromotionModel) Insert(ctx context.Context, data *SmsFlashPromotion) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultSmsFlashPromotionModel) Update(ctx context.Context, data *SmsFlashPromotion) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
