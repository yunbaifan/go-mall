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
	umsIntegrationConsumeSettingModel interface {
		Insert(ctx context.Context, data *UmsIntegrationConsumeSetting) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsIntegrationConsumeSetting, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, keyword *KeywordUmsIntegrationConsumeSettingModel) (*UmsIntegrationConsumeSettingPagination, error)
		Update(ctx context.Context, data *UmsIntegrationConsumeSetting) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsIntegrationConsumeSettingModel struct {
		table   string
		ormConn *gorm.DB
	}

	UmsIntegrationConsumeSetting struct {
		Id                 int64         `json:"id" gorm:"primary_key"`
		DeductionPerAmount sql.NullInt64 `json:"deduction_per_amount"`  // 每一元需要抵扣的积分数量
		MaxPercentPerOrder sql.NullInt64 `json:"max_percent_per_order"` // 每笔订单最高抵用百分比
		UseUnit            sql.NullInt64 `json:"use_unit"`              // 每次使用积分最小单位100
		CouponStatus       sql.NullInt64 `json:"coupon_status"`         // 是否可以和优惠券同用；0->不可以；1->可以
		UmsIntegrationConsumeSettingRelation
	}

	UmsIntegrationConsumeSettingPagination struct {
		Data []*UmsIntegrationConsumeSetting
		UmsIntegrationConsumeSettingPaginationBasic
	}

	UmsIntegrationConsumeSettingPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordUmsIntegrationConsumeSettingModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newUmsIntegrationConsumeSettingModel(db *gorm.DB) *defaultUmsIntegrationConsumeSettingModel {
	return &defaultUmsIntegrationConsumeSettingModel{
		table:   "`ums_integration_consume_setting`",
		ormConn: db,
	}
}

func (UmsIntegrationConsumeSetting) TableName() string {
	return "`ums_integration_consume_setting`"
}

func (m UmsIntegrationConsumeSetting) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultUmsIntegrationConsumeSettingModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&UmsIntegrationConsumeSetting{}, id).Error
}

func (m *defaultUmsIntegrationConsumeSettingModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsIntegrationConsumeSetting, error) {
	var mm UmsIntegrationConsumeSetting
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

func (m *defaultUmsIntegrationConsumeSettingModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&UmsIntegrationConsumeSetting{}).WithContext(ctx)
}

func (m *defaultUmsIntegrationConsumeSettingModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultUmsIntegrationConsumeSettingModel) FindPageListByBuilder(db *gorm.DB, keyword *KeywordUmsIntegrationConsumeSettingModel) (*UmsIntegrationConsumeSettingPagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(UmsIntegrationConsumeSettingPagination)
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

func (m UmsIntegrationConsumeSettingPagination) Pagination() UmsIntegrationConsumeSettingPaginationBasic {
	return UmsIntegrationConsumeSettingPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m UmsIntegrationConsumeSettingPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultUmsIntegrationConsumeSettingModel) Insert(ctx context.Context, data *UmsIntegrationConsumeSetting) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultUmsIntegrationConsumeSettingModel) Update(ctx context.Context, data *UmsIntegrationConsumeSetting) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
