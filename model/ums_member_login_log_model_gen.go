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
	umsMemberLoginLogModel interface {
		Insert(ctx context.Context, data *UmsMemberLoginLog) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsMemberLoginLog, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ctx context.Context, keyword *KeywordUmsMemberLoginLogModel) (*UmsMemberLoginLogPagination, error)
		Update(ctx context.Context, data *UmsMemberLoginLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsMemberLoginLogModel struct {
		table   string
		ormConn *gorm.DB
	}

	UmsMemberLoginLog struct {
		Id         int64          `json:"id" gorm:"primary_key"`
		MemberId   sql.NullInt64  `json:"member_id"`
		CreateTime sql.NullTime   `json:"create_time"`
		Ip         sql.NullString `json:"ip"`
		City       sql.NullString `json:"city"`
		LoginType  sql.NullInt64  `json:"login_type"` // 登录类型：0->PC；1->android;2->ios;3->小程序
		Province   sql.NullString `json:"province"`
		UmsMemberLoginLogRelation
	}

	UmsMemberLoginLogPagination struct {
		Data []*UmsMemberLoginLog
		UmsMemberLoginLogPaginationBasic
	}

	UmsMemberLoginLogPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordUmsMemberLoginLogModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newUmsMemberLoginLogModel(db *gorm.DB) *defaultUmsMemberLoginLogModel {
	return &defaultUmsMemberLoginLogModel{
		table:   "`ums_member_login_log`",
		ormConn: db,
	}
}

func (UmsMemberLoginLog) TableName() string {
	return "`ums_member_login_log`"
}

func (m UmsMemberLoginLog) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultUmsMemberLoginLogModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&UmsMemberLoginLog{}, id).Error
}

func (m *defaultUmsMemberLoginLogModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsMemberLoginLog, error) {
	var mm UmsMemberLoginLog
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

func (m *defaultUmsMemberLoginLogModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&UmsMemberLoginLog{}).WithContext(ctx)
}

func (m *defaultUmsMemberLoginLogModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultUmsMemberLoginLogModel) FindPageListByBuilder(ctx context.Context, keyword *KeywordUmsMemberLoginLogModel) (*UmsMemberLoginLogPagination, error) {
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

	resp := new(UmsMemberLoginLogPagination)
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

func (m UmsMemberLoginLogPagination) Pagination() UmsMemberLoginLogPaginationBasic {
	return UmsMemberLoginLogPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m UmsMemberLoginLogPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultUmsMemberLoginLogModel) Insert(ctx context.Context, data *UmsMemberLoginLog) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultUmsMemberLoginLogModel) Update(ctx context.Context, data *UmsMemberLoginLog) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
