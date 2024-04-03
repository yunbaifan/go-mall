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
	smsFlashPromotionLogModel interface {
		Insert(ctx context.Context, data *SmsFlashPromotionLog) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*SmsFlashPromotionLog, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, page, pageSize int64) (*SmsFlashPromotionLogPagination, error)
		Update(ctx context.Context, data *SmsFlashPromotionLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSmsFlashPromotionLogModel struct {
		table   string
		ormConn *gorm.DB
	}

	SmsFlashPromotionLog struct {
		Id            int64          `json:"id" gorm:"primary_key"`
		MemberId      sql.NullInt64  `json:"member_id"`
		ProductId     sql.NullInt64  `json:"product_id"`
		MemberPhone   sql.NullString `json:"member_phone"`
		ProductName   sql.NullString `json:"product_name"`
		SubscribeTime sql.NullTime   `json:"subscribe_time"` // 会员订阅时间
		SendTime      sql.NullTime   `json:"send_time"`
		SmsFlashPromotionLogRelation
	}

	SmsFlashPromotionLogPagination struct {
		Data []*SmsFlashPromotionLog
		SmsFlashPromotionLogPaginationBasic
	}

	SmsFlashPromotionLogPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}
)

func newSmsFlashPromotionLogModel(db *gorm.DB) *defaultSmsFlashPromotionLogModel {
	return &defaultSmsFlashPromotionLogModel{
		table:   "`sms_flash_promotion_log`",
		ormConn: db,
	}
}

func (SmsFlashPromotionLog) TableName() string {
	return "`sms_flash_promotion_log`"
}

func (m SmsFlashPromotionLog) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultSmsFlashPromotionLogModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&SmsFlashPromotionLog{}, id).Error
}

func (m *defaultSmsFlashPromotionLogModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*SmsFlashPromotionLog, error) {
	var mm SmsFlashPromotionLog
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

func (m *defaultSmsFlashPromotionLogModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&SmsFlashPromotionLog{}).WithContext(ctx)
}

func (m *defaultSmsFlashPromotionLogModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultSmsFlashPromotionLogModel) FindPageListByBuilder(db *gorm.DB, page, pageSize int64) (*SmsFlashPromotionLogPagination, error) {
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(SmsFlashPromotionLogPagination)
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

func (m SmsFlashPromotionLogPagination) Pagination() SmsFlashPromotionLogPaginationBasic {
	return SmsFlashPromotionLogPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m SmsFlashPromotionLogPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultSmsFlashPromotionLogModel) Insert(ctx context.Context, data *SmsFlashPromotionLog) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultSmsFlashPromotionLogModel) Update(ctx context.Context, data *SmsFlashPromotionLog) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}