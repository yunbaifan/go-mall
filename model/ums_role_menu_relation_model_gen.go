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
	umsRoleMenuRelationModel interface {
		Insert(ctx context.Context, data *UmsRoleMenuRelation) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsRoleMenuRelation, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, keyword *KeywordUmsRoleMenuRelationModel) (*UmsRoleMenuRelationPagination, error)
		Update(ctx context.Context, data *UmsRoleMenuRelation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsRoleMenuRelationModel struct {
		table   string
		ormConn *gorm.DB
	}

	UmsRoleMenuRelation struct {
		Id     int64         `json:"id" gorm:"primary_key"`
		RoleId sql.NullInt64 `json:"role_id"` // 角色ID
		MenuId sql.NullInt64 `json:"menu_id"` // 菜单ID
		UmsRoleMenuRelationRelation
	}

	UmsRoleMenuRelationPagination struct {
		Data []*UmsRoleMenuRelation
		UmsRoleMenuRelationPaginationBasic
	}

	UmsRoleMenuRelationPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordUmsRoleMenuRelationModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newUmsRoleMenuRelationModel(db *gorm.DB) *defaultUmsRoleMenuRelationModel {
	return &defaultUmsRoleMenuRelationModel{
		table:   "`ums_role_menu_relation`",
		ormConn: db,
	}
}

func (UmsRoleMenuRelation) TableName() string {
	return "`ums_role_menu_relation`"
}

func (m UmsRoleMenuRelation) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultUmsRoleMenuRelationModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&UmsRoleMenuRelation{}, id).Error
}

func (m *defaultUmsRoleMenuRelationModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsRoleMenuRelation, error) {
	var mm UmsRoleMenuRelation
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

func (m *defaultUmsRoleMenuRelationModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&UmsRoleMenuRelation{}).WithContext(ctx)
}

func (m *defaultUmsRoleMenuRelationModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultUmsRoleMenuRelationModel) FindPageListByBuilder(db *gorm.DB, keyword *KeywordUmsRoleMenuRelationModel) (*UmsRoleMenuRelationPagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(UmsRoleMenuRelationPagination)
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

func (m UmsRoleMenuRelationPagination) Pagination() UmsRoleMenuRelationPaginationBasic {
	return UmsRoleMenuRelationPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m UmsRoleMenuRelationPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultUmsRoleMenuRelationModel) Insert(ctx context.Context, data *UmsRoleMenuRelation) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultUmsRoleMenuRelationModel) Update(ctx context.Context, data *UmsRoleMenuRelation) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
