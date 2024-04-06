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
	umsRoleModel interface {
		Insert(ctx context.Context, data *UmsRole) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsRole, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, keyword *KeywordUmsRoleModel) (*UmsRolePagination, error)
		Update(ctx context.Context, data *UmsRole) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsRoleModel struct {
		table   string
		ormConn *gorm.DB
	}

	UmsRole struct {
		Id          int64          `json:"id" gorm:"primary_key"`
		Name        sql.NullString `json:"name"`        // 名称
		Description sql.NullString `json:"description"` // 描述
		AdminCount  sql.NullInt64  `json:"admin_count"` // 后台用户数量
		CreateTime  sql.NullTime   `json:"create_time"` // 创建时间
		Status      int64          `json:"status"`      // 启用状态：0->禁用；1->启用
		Sort        int64          `json:"sort"`
		UmsRoleRelation
	}

	UmsRolePagination struct {
		Data []*UmsRole
		UmsRolePaginationBasic
	}

	UmsRolePaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordUmsRoleModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newUmsRoleModel(db *gorm.DB) *defaultUmsRoleModel {
	return &defaultUmsRoleModel{
		table:   "`ums_role`",
		ormConn: db,
	}
}

func (UmsRole) TableName() string {
	return "`ums_role`"
}

func (m UmsRole) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultUmsRoleModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&UmsRole{}, id).Error
}

func (m *defaultUmsRoleModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsRole, error) {
	var mm UmsRole
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

func (m *defaultUmsRoleModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&UmsRole{}).WithContext(ctx)
}

func (m *defaultUmsRoleModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultUmsRoleModel) FindPageListByBuilder(db *gorm.DB, keyword *KeywordUmsRoleModel) (*UmsRolePagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(UmsRolePagination)
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

func (m UmsRolePagination) Pagination() UmsRolePaginationBasic {
	return UmsRolePaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m UmsRolePagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultUmsRoleModel) Insert(ctx context.Context, data *UmsRole) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultUmsRoleModel) Update(ctx context.Context, data *UmsRole) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
