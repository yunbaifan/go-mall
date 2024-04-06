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
	pmsMemberPriceModel interface {
		Insert(ctx context.Context, data *PmsMemberPrice) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsMemberPrice, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, keyword *KeywordPmsMemberPriceModel) (*PmsMemberPricePagination, error)
		Update(ctx context.Context, data *PmsMemberPrice) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsMemberPriceModel struct {
		table   string
		ormConn *gorm.DB
	}

	PmsMemberPrice struct {
		Id              int64           `json:"id" gorm:"primary_key"`
		ProductId       sql.NullInt64   `json:"product_id"`
		MemberLevelId   sql.NullInt64   `json:"member_level_id"`
		MemberPrice     sql.NullFloat64 `json:"member_price"` // 会员价格
		MemberLevelName sql.NullString  `json:"member_level_name"`
		PmsMemberPriceRelation
	}

	PmsMemberPricePagination struct {
		Data []*PmsMemberPrice
		PmsMemberPricePaginationBasic
	}

	PmsMemberPricePaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordPmsMemberPriceModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newPmsMemberPriceModel(db *gorm.DB) *defaultPmsMemberPriceModel {
	return &defaultPmsMemberPriceModel{
		table:   "`pms_member_price`",
		ormConn: db,
	}
}

func (PmsMemberPrice) TableName() string {
	return "`pms_member_price`"
}

func (m PmsMemberPrice) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultPmsMemberPriceModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&PmsMemberPrice{}, id).Error
}

func (m *defaultPmsMemberPriceModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsMemberPrice, error) {
	var mm PmsMemberPrice
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

func (m *defaultPmsMemberPriceModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&PmsMemberPrice{}).WithContext(ctx)
}

func (m *defaultPmsMemberPriceModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultPmsMemberPriceModel) FindPageListByBuilder(db *gorm.DB, keyword *KeywordPmsMemberPriceModel) (*PmsMemberPricePagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(PmsMemberPricePagination)
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

func (m PmsMemberPricePagination) Pagination() PmsMemberPricePaginationBasic {
	return PmsMemberPricePaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m PmsMemberPricePagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultPmsMemberPriceModel) Insert(ctx context.Context, data *PmsMemberPrice) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultPmsMemberPriceModel) Update(ctx context.Context, data *PmsMemberPrice) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
