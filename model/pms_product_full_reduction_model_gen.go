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
	pmsProductFullReductionModel interface {
		Insert(ctx context.Context, data *PmsProductFullReduction) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsProductFullReduction, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ctx context.Context, keyword *KeywordPmsProductFullReductionModel) (*PmsProductFullReductionPagination, error)
		Update(ctx context.Context, data *PmsProductFullReduction) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsProductFullReductionModel struct {
		table   string
		ormConn *gorm.DB
	}

	PmsProductFullReduction struct {
		Id          int64           `json:"id" gorm:"primary_key"`
		ProductId   sql.NullInt64   `json:"product_id"`
		FullPrice   sql.NullFloat64 `json:"full_price"`
		ReducePrice sql.NullFloat64 `json:"reduce_price"`
		PmsProductFullReductionRelation
	}

	PmsProductFullReductionPagination struct {
		Data []*PmsProductFullReduction
		PmsProductFullReductionPaginationBasic
	}

	PmsProductFullReductionPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordPmsProductFullReductionModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newPmsProductFullReductionModel(db *gorm.DB) *defaultPmsProductFullReductionModel {
	return &defaultPmsProductFullReductionModel{
		table:   "`pms_product_full_reduction`",
		ormConn: db,
	}
}

func (PmsProductFullReduction) TableName() string {
	return "`pms_product_full_reduction`"
}

func (m PmsProductFullReduction) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultPmsProductFullReductionModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&PmsProductFullReduction{}, id).Error
}

func (m *defaultPmsProductFullReductionModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsProductFullReduction, error) {
	var mm PmsProductFullReduction
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

func (m *defaultPmsProductFullReductionModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&PmsProductFullReduction{}).WithContext(ctx)
}

func (m *defaultPmsProductFullReductionModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultPmsProductFullReductionModel) FindPageListByBuilder(ctx context.Context, keyword *KeywordPmsProductFullReductionModel) (*PmsProductFullReductionPagination, error) {
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

	resp := new(PmsProductFullReductionPagination)
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

func (m PmsProductFullReductionPagination) Pagination() PmsProductFullReductionPaginationBasic {
	return PmsProductFullReductionPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m PmsProductFullReductionPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultPmsProductFullReductionModel) Insert(ctx context.Context, data *PmsProductFullReduction) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultPmsProductFullReductionModel) Update(ctx context.Context, data *PmsProductFullReduction) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
