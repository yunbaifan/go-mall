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
	smsHomeRecommendSubjectModel interface {
		Insert(ctx context.Context, data *SmsHomeRecommendSubject) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*SmsHomeRecommendSubject, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, keyword *KeywordSmsHomeRecommendSubjectModel) (*SmsHomeRecommendSubjectPagination, error)
		Update(ctx context.Context, data *SmsHomeRecommendSubject) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSmsHomeRecommendSubjectModel struct {
		table   string
		ormConn *gorm.DB
	}

	SmsHomeRecommendSubject struct {
		Id              int64          `json:"id" gorm:"primary_key"`
		SubjectId       sql.NullInt64  `json:"subject_id"`
		SubjectName     sql.NullString `json:"subject_name"`
		RecommendStatus sql.NullInt64  `json:"recommend_status"`
		Sort            sql.NullInt64  `json:"sort"`
		SmsHomeRecommendSubjectRelation
	}

	SmsHomeRecommendSubjectPagination struct {
		Data []*SmsHomeRecommendSubject
		SmsHomeRecommendSubjectPaginationBasic
	}

	SmsHomeRecommendSubjectPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordSmsHomeRecommendSubjectModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newSmsHomeRecommendSubjectModel(db *gorm.DB) *defaultSmsHomeRecommendSubjectModel {
	return &defaultSmsHomeRecommendSubjectModel{
		table:   "`sms_home_recommend_subject`",
		ormConn: db,
	}
}

func (SmsHomeRecommendSubject) TableName() string {
	return "`sms_home_recommend_subject`"
}

func (m SmsHomeRecommendSubject) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultSmsHomeRecommendSubjectModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&SmsHomeRecommendSubject{}, id).Error
}

func (m *defaultSmsHomeRecommendSubjectModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*SmsHomeRecommendSubject, error) {
	var mm SmsHomeRecommendSubject
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

func (m *defaultSmsHomeRecommendSubjectModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&SmsHomeRecommendSubject{}).WithContext(ctx)
}

func (m *defaultSmsHomeRecommendSubjectModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultSmsHomeRecommendSubjectModel) FindPageListByBuilder(db *gorm.DB, keyword *KeywordSmsHomeRecommendSubjectModel) (*SmsHomeRecommendSubjectPagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(SmsHomeRecommendSubjectPagination)
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

func (m SmsHomeRecommendSubjectPagination) Pagination() SmsHomeRecommendSubjectPaginationBasic {
	return SmsHomeRecommendSubjectPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m SmsHomeRecommendSubjectPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultSmsHomeRecommendSubjectModel) Insert(ctx context.Context, data *SmsHomeRecommendSubject) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultSmsHomeRecommendSubjectModel) Update(ctx context.Context, data *SmsHomeRecommendSubject) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
