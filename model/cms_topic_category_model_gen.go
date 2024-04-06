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
	cmsTopicCategoryModel interface {
		Insert(ctx context.Context, data *CmsTopicCategory) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*CmsTopicCategory, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ctx context.Context, keyword *KeywordCmsTopicCategoryModel) (*CmsTopicCategoryPagination, error)
		Update(ctx context.Context, data *CmsTopicCategory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCmsTopicCategoryModel struct {
		table   string
		ormConn *gorm.DB
	}

	CmsTopicCategory struct {
		Id           int64          `json:"id" gorm:"primary_key"`
		Name         sql.NullString `json:"name"`
		Icon         sql.NullString `json:"icon"`          // 分类图标
		SubjectCount sql.NullInt64  `json:"subject_count"` // 专题数量
		ShowStatus   sql.NullInt64  `json:"show_status"`
		Sort         sql.NullInt64  `json:"sort"`
		CmsTopicCategoryRelation
	}

	CmsTopicCategoryPagination struct {
		Data []*CmsTopicCategory
		CmsTopicCategoryPaginationBasic
	}

	CmsTopicCategoryPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordCmsTopicCategoryModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newCmsTopicCategoryModel(db *gorm.DB) *defaultCmsTopicCategoryModel {
	return &defaultCmsTopicCategoryModel{
		table:   "`cms_topic_category`",
		ormConn: db,
	}
}

func (CmsTopicCategory) TableName() string {
	return "`cms_topic_category`"
}

func (m CmsTopicCategory) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultCmsTopicCategoryModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&CmsTopicCategory{}, id).Error
}

func (m *defaultCmsTopicCategoryModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*CmsTopicCategory, error) {
	var mm CmsTopicCategory
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

func (m *defaultCmsTopicCategoryModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&CmsTopicCategory{}).WithContext(ctx)
}

func (m *defaultCmsTopicCategoryModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultCmsTopicCategoryModel) FindPageListByBuilder(ctx context.Context, keyword *KeywordCmsTopicCategoryModel) (*CmsTopicCategoryPagination, error) {
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

	resp := new(CmsTopicCategoryPagination)
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

func (m CmsTopicCategoryPagination) Pagination() CmsTopicCategoryPaginationBasic {
	return CmsTopicCategoryPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m CmsTopicCategoryPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultCmsTopicCategoryModel) Insert(ctx context.Context, data *CmsTopicCategory) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultCmsTopicCategoryModel) Update(ctx context.Context, data *CmsTopicCategory) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
