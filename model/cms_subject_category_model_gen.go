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
	cmsSubjectCategoryModel interface {
		Insert(ctx context.Context, data *CmsSubjectCategory) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*CmsSubjectCategory, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ctx context.Context, keyword *KeywordCmsSubjectCategoryModel) (*CmsSubjectCategoryPagination, error)
		Update(ctx context.Context, data *CmsSubjectCategory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCmsSubjectCategoryModel struct {
		table   string
		ormConn *gorm.DB
	}

	CmsSubjectCategory struct {
		Id           int64          `json:"id" gorm:"primary_key"`
		Name         sql.NullString `json:"name"`
		Icon         sql.NullString `json:"icon"`          // 分类图标
		SubjectCount sql.NullInt64  `json:"subject_count"` // 专题数量
		ShowStatus   sql.NullInt64  `json:"show_status"`
		Sort         sql.NullInt64  `json:"sort"`
		CmsSubjectCategoryRelation
	}

	CmsSubjectCategoryPagination struct {
		Data []*CmsSubjectCategory
		CmsSubjectCategoryPaginationBasic
	}

	CmsSubjectCategoryPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordCmsSubjectCategoryModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newCmsSubjectCategoryModel(db *gorm.DB) *defaultCmsSubjectCategoryModel {
	return &defaultCmsSubjectCategoryModel{
		table:   "`cms_subject_category`",
		ormConn: db,
	}
}

func (CmsSubjectCategory) TableName() string {
	return "`cms_subject_category`"
}

func (m CmsSubjectCategory) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultCmsSubjectCategoryModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&CmsSubjectCategory{}, id).Error
}

func (m *defaultCmsSubjectCategoryModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*CmsSubjectCategory, error) {
	var mm CmsSubjectCategory
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

func (m *defaultCmsSubjectCategoryModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&CmsSubjectCategory{}).WithContext(ctx)
}

func (m *defaultCmsSubjectCategoryModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultCmsSubjectCategoryModel) FindPageListByBuilder(ctx context.Context, keyword *KeywordCmsSubjectCategoryModel) (*CmsSubjectCategoryPagination, error) {
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

	resp := new(CmsSubjectCategoryPagination)
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

func (m CmsSubjectCategoryPagination) Pagination() CmsSubjectCategoryPaginationBasic {
	return CmsSubjectCategoryPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m CmsSubjectCategoryPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultCmsSubjectCategoryModel) Insert(ctx context.Context, data *CmsSubjectCategory) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultCmsSubjectCategoryModel) Update(ctx context.Context, data *CmsSubjectCategory) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
