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
	cmsHelpCategoryModel interface {
		Insert(ctx context.Context, data *CmsHelpCategory) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*CmsHelpCategory, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, keyword *KeywordCmsHelpCategoryModel) (*CmsHelpCategoryPagination, error)
		Update(ctx context.Context, data *CmsHelpCategory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCmsHelpCategoryModel struct {
		table   string
		ormConn *gorm.DB
	}

	CmsHelpCategory struct {
		Id         int64          `json:"id" gorm:"primary_key"`
		Name       sql.NullString `json:"name"`
		Icon       sql.NullString `json:"icon"`       // 分类图标
		HelpCount  sql.NullInt64  `json:"help_count"` // 专题数量
		ShowStatus sql.NullInt64  `json:"show_status"`
		Sort       sql.NullInt64  `json:"sort"`
		CmsHelpCategoryRelation
	}

	CmsHelpCategoryPagination struct {
		Data []*CmsHelpCategory
		CmsHelpCategoryPaginationBasic
	}

	CmsHelpCategoryPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordCmsHelpCategoryModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newCmsHelpCategoryModel(db *gorm.DB) *defaultCmsHelpCategoryModel {
	return &defaultCmsHelpCategoryModel{
		table:   "`cms_help_category`",
		ormConn: db,
	}
}

func (CmsHelpCategory) TableName() string {
	return "`cms_help_category`"
}

func (m CmsHelpCategory) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultCmsHelpCategoryModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&CmsHelpCategory{}, id).Error
}

func (m *defaultCmsHelpCategoryModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*CmsHelpCategory, error) {
	var mm CmsHelpCategory
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

func (m *defaultCmsHelpCategoryModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&CmsHelpCategory{}).WithContext(ctx)
}

func (m *defaultCmsHelpCategoryModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultCmsHelpCategoryModel) FindPageListByBuilder(db *gorm.DB, keyword *KeywordCmsHelpCategoryModel) (*CmsHelpCategoryPagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(CmsHelpCategoryPagination)
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

func (m CmsHelpCategoryPagination) Pagination() CmsHelpCategoryPaginationBasic {
	return CmsHelpCategoryPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m CmsHelpCategoryPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultCmsHelpCategoryModel) Insert(ctx context.Context, data *CmsHelpCategory) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultCmsHelpCategoryModel) Update(ctx context.Context, data *CmsHelpCategory) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
