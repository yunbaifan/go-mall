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
	cmsSubjectCommentModel interface {
		Insert(ctx context.Context, data *CmsSubjectComment) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*CmsSubjectComment, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ctx context.Context, keyword *KeywordCmsSubjectCommentModel) (*CmsSubjectCommentPagination, error)
		Update(ctx context.Context, data *CmsSubjectComment) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCmsSubjectCommentModel struct {
		table   string
		ormConn *gorm.DB
	}

	CmsSubjectComment struct {
		Id             int64          `json:"id" gorm:"primary_key"`
		SubjectId      sql.NullInt64  `json:"subject_id"`
		MemberNickName sql.NullString `json:"member_nick_name"`
		MemberIcon     sql.NullString `json:"member_icon"`
		Content        sql.NullString `json:"content"`
		CreateTime     sql.NullTime   `json:"create_time"`
		ShowStatus     sql.NullInt64  `json:"show_status"`
		CmsSubjectCommentRelation
	}

	CmsSubjectCommentPagination struct {
		Data []*CmsSubjectComment
		CmsSubjectCommentPaginationBasic
	}

	CmsSubjectCommentPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordCmsSubjectCommentModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newCmsSubjectCommentModel(db *gorm.DB) *defaultCmsSubjectCommentModel {
	return &defaultCmsSubjectCommentModel{
		table:   "`cms_subject_comment`",
		ormConn: db,
	}
}

func (CmsSubjectComment) TableName() string {
	return "`cms_subject_comment`"
}

func (m CmsSubjectComment) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultCmsSubjectCommentModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&CmsSubjectComment{}, id).Error
}

func (m *defaultCmsSubjectCommentModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*CmsSubjectComment, error) {
	var mm CmsSubjectComment
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

func (m *defaultCmsSubjectCommentModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&CmsSubjectComment{}).WithContext(ctx)
}

func (m *defaultCmsSubjectCommentModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultCmsSubjectCommentModel) FindPageListByBuilder(ctx context.Context, keyword *KeywordCmsSubjectCommentModel) (*CmsSubjectCommentPagination, error) {
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

	resp := new(CmsSubjectCommentPagination)
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

func (m CmsSubjectCommentPagination) Pagination() CmsSubjectCommentPaginationBasic {
	return CmsSubjectCommentPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m CmsSubjectCommentPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultCmsSubjectCommentModel) Insert(ctx context.Context, data *CmsSubjectComment) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultCmsSubjectCommentModel) Update(ctx context.Context, data *CmsSubjectComment) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
