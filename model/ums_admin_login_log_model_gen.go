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
	umsAdminLoginLogModel interface {
		Insert(ctx context.Context, data *UmsAdminLoginLog) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsAdminLoginLog, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, keyword *KeywordUmsAdminLoginLogModel) (*UmsAdminLoginLogPagination, error)
		Update(ctx context.Context, data *UmsAdminLoginLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsAdminLoginLogModel struct {
		table   string
		ormConn *gorm.DB
	}

	UmsAdminLoginLog struct {
		Id         int64          `json:"id" gorm:"primary_key"`
		AdminId    sql.NullInt64  `json:"admin_id"`
		CreateTime sql.NullTime   `json:"create_time"`
		Ip         sql.NullString `json:"ip"`
		Address    sql.NullString `json:"address"`
		UserAgent  sql.NullString `json:"user_agent"` // 浏览器登录类型
		UmsAdminLoginLogRelation
	}

	UmsAdminLoginLogPagination struct {
		Data []*UmsAdminLoginLog
		UmsAdminLoginLogPaginationBasic
	}

	UmsAdminLoginLogPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordUmsAdminLoginLogModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newUmsAdminLoginLogModel(db *gorm.DB) *defaultUmsAdminLoginLogModel {
	return &defaultUmsAdminLoginLogModel{
		table:   "`ums_admin_login_log`",
		ormConn: db,
	}
}

func (UmsAdminLoginLog) TableName() string {
	return "`ums_admin_login_log`"
}

func (m UmsAdminLoginLog) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultUmsAdminLoginLogModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&UmsAdminLoginLog{}, id).Error
}

func (m *defaultUmsAdminLoginLogModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsAdminLoginLog, error) {
	var mm UmsAdminLoginLog
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

func (m *defaultUmsAdminLoginLogModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&UmsAdminLoginLog{}).WithContext(ctx)
}

func (m *defaultUmsAdminLoginLogModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultUmsAdminLoginLogModel) FindPageListByBuilder(db *gorm.DB, keyword *KeywordUmsAdminLoginLogModel) (*UmsAdminLoginLogPagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(UmsAdminLoginLogPagination)
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

func (m UmsAdminLoginLogPagination) Pagination() UmsAdminLoginLogPaginationBasic {
	return UmsAdminLoginLogPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m UmsAdminLoginLogPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultUmsAdminLoginLogModel) Insert(ctx context.Context, data *UmsAdminLoginLog) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultUmsAdminLoginLogModel) Update(ctx context.Context, data *UmsAdminLoginLog) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
