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
	smsFlashPromotionProductRelationModel interface {
		Insert(ctx context.Context, data *SmsFlashPromotionProductRelation) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*SmsFlashPromotionProductRelation, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ctx context.Context, keyword *KeywordSmsFlashPromotionProductRelationModel) (*SmsFlashPromotionProductRelationPagination, error)
		Update(ctx context.Context, data *SmsFlashPromotionProductRelation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSmsFlashPromotionProductRelationModel struct {
		table   string
		ormConn *gorm.DB
	}

	SmsFlashPromotionProductRelation struct {
		Id                      int64           `json:"id" gorm:"primary_key"` // 编号
		FlashPromotionId        sql.NullInt64   `json:"flash_promotion_id"`
		FlashPromotionSessionId sql.NullInt64   `json:"flash_promotion_session_id"` // 编号
		ProductId               sql.NullInt64   `json:"product_id"`
		FlashPromotionPrice     sql.NullFloat64 `json:"flash_promotion_price"` // 限时购价格
		FlashPromotionCount     sql.NullInt64   `json:"flash_promotion_count"` // 限时购数量
		FlashPromotionLimit     sql.NullInt64   `json:"flash_promotion_limit"` // 每人限购数量
		Sort                    sql.NullInt64   `json:"sort"`                  // 排序
		SmsFlashPromotionProductRelationRelation
	}

	SmsFlashPromotionProductRelationPagination struct {
		Data []*SmsFlashPromotionProductRelation
		SmsFlashPromotionProductRelationPaginationBasic
	}

	SmsFlashPromotionProductRelationPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordSmsFlashPromotionProductRelationModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newSmsFlashPromotionProductRelationModel(db *gorm.DB) *defaultSmsFlashPromotionProductRelationModel {
	return &defaultSmsFlashPromotionProductRelationModel{
		table:   "`sms_flash_promotion_product_relation`",
		ormConn: db,
	}
}

func (SmsFlashPromotionProductRelation) TableName() string {
	return "`sms_flash_promotion_product_relation`"
}

func (m SmsFlashPromotionProductRelation) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultSmsFlashPromotionProductRelationModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&SmsFlashPromotionProductRelation{}, id).Error
}

func (m *defaultSmsFlashPromotionProductRelationModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*SmsFlashPromotionProductRelation, error) {
	var mm SmsFlashPromotionProductRelation
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

func (m *defaultSmsFlashPromotionProductRelationModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&SmsFlashPromotionProductRelation{}).WithContext(ctx)
}

func (m *defaultSmsFlashPromotionProductRelationModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultSmsFlashPromotionProductRelationModel) FindPageListByBuilder(ctx context.Context, keyword *KeywordSmsFlashPromotionProductRelationModel) (*SmsFlashPromotionProductRelationPagination, error) {
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

	resp := new(SmsFlashPromotionProductRelationPagination)
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

func (m SmsFlashPromotionProductRelationPagination) Pagination() SmsFlashPromotionProductRelationPaginationBasic {
	return SmsFlashPromotionProductRelationPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m SmsFlashPromotionProductRelationPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultSmsFlashPromotionProductRelationModel) Insert(ctx context.Context, data *SmsFlashPromotionProductRelation) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultSmsFlashPromotionProductRelationModel) Update(ctx context.Context, data *SmsFlashPromotionProductRelation) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
