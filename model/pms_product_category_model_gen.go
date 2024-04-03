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
	pmsProductCategoryModel interface {
		Insert(ctx context.Context, data *PmsProductCategory) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsProductCategory, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, page, pageSize int64) (*PmsProductCategoryPagination, error)
		Update(ctx context.Context, data *PmsProductCategory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsProductCategoryModel struct {
		table   string
		ormConn *gorm.DB
	}

	PmsProductCategory struct {
		Id           int64          `json:"id" gorm:"primary_key"`
		ParentId     sql.NullInt64  `json:"parent_id"` // 上机分类的编号：0表示一级分类
		Name         sql.NullString `json:"name"`
		Level        sql.NullInt64  `json:"level"` // 分类级别：0->1级；1->2级
		ProductCount sql.NullInt64  `json:"product_count"`
		ProductUnit  sql.NullString `json:"product_unit"`
		NavStatus    sql.NullInt64  `json:"nav_status"`  // 是否显示在导航栏：0->不显示；1->显示
		ShowStatus   sql.NullInt64  `json:"show_status"` // 显示状态：0->不显示；1->显示
		Sort         sql.NullInt64  `json:"sort"`
		Icon         sql.NullString `json:"icon"` // 图标
		Keywords     sql.NullString `json:"keywords"`
		Description  sql.NullString `json:"description"` // 描述
		PmsProductCategoryRelation
	}

	PmsProductCategoryPagination struct {
		Data []*PmsProductCategory
		PmsProductCategoryPaginationBasic
	}

	PmsProductCategoryPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}
)

func newPmsProductCategoryModel(db *gorm.DB) *defaultPmsProductCategoryModel {
	return &defaultPmsProductCategoryModel{
		table:   "`pms_product_category`",
		ormConn: db,
	}
}

func (PmsProductCategory) TableName() string {
	return "`pms_product_category`"
}

func (m PmsProductCategory) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultPmsProductCategoryModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&PmsProductCategory{}, id).Error
}

func (m *defaultPmsProductCategoryModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsProductCategory, error) {
	var mm PmsProductCategory
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

func (m *defaultPmsProductCategoryModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&PmsProductCategory{}).WithContext(ctx)
}

func (m *defaultPmsProductCategoryModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultPmsProductCategoryModel) FindPageListByBuilder(db *gorm.DB, page, pageSize int64) (*PmsProductCategoryPagination, error) {
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(PmsProductCategoryPagination)
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

func (m PmsProductCategoryPagination) Pagination() PmsProductCategoryPaginationBasic {
	return PmsProductCategoryPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m PmsProductCategoryPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultPmsProductCategoryModel) Insert(ctx context.Context, data *PmsProductCategory) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultPmsProductCategoryModel) Update(ctx context.Context, data *PmsProductCategory) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
