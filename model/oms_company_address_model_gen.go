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
	omsCompanyAddressModel interface {
		Insert(ctx context.Context, data *OmsCompanyAddress) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*OmsCompanyAddress, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ctx context.Context, keyword *KeywordOmsCompanyAddressModel) (*OmsCompanyAddressPagination, error)
		Update(ctx context.Context, data *OmsCompanyAddress) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOmsCompanyAddressModel struct {
		table   string
		ormConn *gorm.DB
	}

	OmsCompanyAddress struct {
		Id            int64          `json:"id" gorm:"primary_key"`
		AddressName   sql.NullString `json:"address_name"`   // 地址名称
		SendStatus    sql.NullInt64  `json:"send_status"`    // 默认发货地址：0->否；1->是
		ReceiveStatus sql.NullInt64  `json:"receive_status"` // 是否默认收货地址：0->否；1->是
		Name          sql.NullString `json:"name"`           // 收发货人姓名
		Phone         sql.NullString `json:"phone"`          // 收货人电话
		Province      sql.NullString `json:"province"`       // 省/直辖市
		City          sql.NullString `json:"city"`           // 市
		Region        sql.NullString `json:"region"`         // 区
		DetailAddress sql.NullString `json:"detail_address"` // 详细地址
		OmsCompanyAddressRelation
	}

	OmsCompanyAddressPagination struct {
		Data []*OmsCompanyAddress
		OmsCompanyAddressPaginationBasic
	}

	OmsCompanyAddressPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordOmsCompanyAddressModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newOmsCompanyAddressModel(db *gorm.DB) *defaultOmsCompanyAddressModel {
	return &defaultOmsCompanyAddressModel{
		table:   "`oms_company_address`",
		ormConn: db,
	}
}

func (OmsCompanyAddress) TableName() string {
	return "`oms_company_address`"
}

func (m OmsCompanyAddress) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultOmsCompanyAddressModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&OmsCompanyAddress{}, id).Error
}

func (m *defaultOmsCompanyAddressModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*OmsCompanyAddress, error) {
	var mm OmsCompanyAddress
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

func (m *defaultOmsCompanyAddressModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&OmsCompanyAddress{}).WithContext(ctx)
}

func (m *defaultOmsCompanyAddressModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultOmsCompanyAddressModel) FindPageListByBuilder(ctx context.Context, keyword *KeywordOmsCompanyAddressModel) (*OmsCompanyAddressPagination, error) {
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

	resp := new(OmsCompanyAddressPagination)
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

func (m OmsCompanyAddressPagination) Pagination() OmsCompanyAddressPaginationBasic {
	return OmsCompanyAddressPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m OmsCompanyAddressPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultOmsCompanyAddressModel) Insert(ctx context.Context, data *OmsCompanyAddress) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultOmsCompanyAddressModel) Update(ctx context.Context, data *OmsCompanyAddress) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
