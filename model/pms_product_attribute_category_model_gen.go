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
	pmsProductAttributeCategoryModel interface {
		Insert(ctx context.Context, data *PmsProductAttributeCategory) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsProductAttributeCategory, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ctx context.Context, keyword *KeywordPmsProductAttributeCategoryModel) (*PmsProductAttributeCategoryPagination, error)
		Update(ctx context.Context, data *PmsProductAttributeCategory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsProductAttributeCategoryModel struct {
		table   string
		ormConn *gorm.DB
	}

	PmsProductAttributeCategory struct {
		Id             int64          `json:"id" gorm:"primary_key"`
		Name           sql.NullString `json:"name"`
		AttributeCount int64          `json:"attribute_count"` // 属性数量
		ParamCount     int64          `json:"param_count"`     // 参数数量
		PmsProductAttributeCategoryRelation
	}

	PmsProductAttributeCategoryPagination struct {
		Data []*PmsProductAttributeCategory
		PmsProductAttributeCategoryPaginationBasic
	}

	PmsProductAttributeCategoryPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordPmsProductAttributeCategoryModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newPmsProductAttributeCategoryModel(db *gorm.DB) *defaultPmsProductAttributeCategoryModel {
	return &defaultPmsProductAttributeCategoryModel{
		table:   "`pms_product_attribute_category`",
		ormConn: db,
	}
}

func (PmsProductAttributeCategory) TableName() string {
	return "`pms_product_attribute_category`"
}

func (m PmsProductAttributeCategory) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultPmsProductAttributeCategoryModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&PmsProductAttributeCategory{}, id).Error
}

func (m *defaultPmsProductAttributeCategoryModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsProductAttributeCategory, error) {
	var mm PmsProductAttributeCategory
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

func (m *defaultPmsProductAttributeCategoryModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&PmsProductAttributeCategory{}).WithContext(ctx)
}

func (m *defaultPmsProductAttributeCategoryModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultPmsProductAttributeCategoryModel) FindPageListByBuilder(ctx context.Context, keyword *KeywordPmsProductAttributeCategoryModel) (*PmsProductAttributeCategoryPagination, error) {
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

	resp := new(PmsProductAttributeCategoryPagination)
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

func (m PmsProductAttributeCategoryPagination) Pagination() PmsProductAttributeCategoryPaginationBasic {
	return PmsProductAttributeCategoryPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m PmsProductAttributeCategoryPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultPmsProductAttributeCategoryModel) Insert(ctx context.Context, data *PmsProductAttributeCategory) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultPmsProductAttributeCategoryModel) Update(ctx context.Context, data *PmsProductAttributeCategory) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
