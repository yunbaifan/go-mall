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
	umsRolePermissionRelationModel interface {
		Insert(ctx context.Context, data *UmsRolePermissionRelation) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsRolePermissionRelation, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, keyword *KeywordUmsRolePermissionRelationModel) (*UmsRolePermissionRelationPagination, error)
		Update(ctx context.Context, data *UmsRolePermissionRelation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsRolePermissionRelationModel struct {
		table   string
		ormConn *gorm.DB
	}

	UmsRolePermissionRelation struct {
		Id           int64         `json:"id" gorm:"primary_key"`
		RoleId       sql.NullInt64 `json:"role_id"`
		PermissionId sql.NullInt64 `json:"permission_id"`
		UmsRolePermissionRelationRelation
	}

	UmsRolePermissionRelationPagination struct {
		Data []*UmsRolePermissionRelation
		UmsRolePermissionRelationPaginationBasic
	}

	UmsRolePermissionRelationPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordUmsRolePermissionRelationModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newUmsRolePermissionRelationModel(db *gorm.DB) *defaultUmsRolePermissionRelationModel {
	return &defaultUmsRolePermissionRelationModel{
		table:   "`ums_role_permission_relation`",
		ormConn: db,
	}
}

func (UmsRolePermissionRelation) TableName() string {
	return "`ums_role_permission_relation`"
}

func (m UmsRolePermissionRelation) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultUmsRolePermissionRelationModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&UmsRolePermissionRelation{}, id).Error
}

func (m *defaultUmsRolePermissionRelationModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsRolePermissionRelation, error) {
	var mm UmsRolePermissionRelation
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

func (m *defaultUmsRolePermissionRelationModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&UmsRolePermissionRelation{}).WithContext(ctx)
}

func (m *defaultUmsRolePermissionRelationModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultUmsRolePermissionRelationModel) FindPageListByBuilder(db *gorm.DB, keyword *KeywordUmsRolePermissionRelationModel) (*UmsRolePermissionRelationPagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(UmsRolePermissionRelationPagination)
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

func (m UmsRolePermissionRelationPagination) Pagination() UmsRolePermissionRelationPaginationBasic {
	return UmsRolePermissionRelationPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m UmsRolePermissionRelationPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultUmsRolePermissionRelationModel) Insert(ctx context.Context, data *UmsRolePermissionRelation) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultUmsRolePermissionRelationModel) Update(ctx context.Context, data *UmsRolePermissionRelation) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
