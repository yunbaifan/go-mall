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
	umsMemberMemberTagRelationModel interface {
		Insert(ctx context.Context, data *UmsMemberMemberTagRelation) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsMemberMemberTagRelation, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, page, pageSize int64) (*UmsMemberMemberTagRelationPagination, error)
		Update(ctx context.Context, data *UmsMemberMemberTagRelation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsMemberMemberTagRelationModel struct {
		table   string
		ormConn *gorm.DB
	}

	UmsMemberMemberTagRelation struct {
		Id       int64         `json:"id" gorm:"primary_key"`
		MemberId sql.NullInt64 `json:"member_id"`
		TagId    sql.NullInt64 `json:"tag_id"`
		UmsMemberMemberTagRelationRelation
	}

	UmsMemberMemberTagRelationPagination struct {
		Data []*UmsMemberMemberTagRelation
		UmsMemberMemberTagRelationPaginationBasic
	}

	UmsMemberMemberTagRelationPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}
)

func newUmsMemberMemberTagRelationModel(db *gorm.DB) *defaultUmsMemberMemberTagRelationModel {
	return &defaultUmsMemberMemberTagRelationModel{
		table:   "`ums_member_member_tag_relation`",
		ormConn: db,
	}
}

func (UmsMemberMemberTagRelation) TableName() string {
	return "`ums_member_member_tag_relation`"
}

func (m UmsMemberMemberTagRelation) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultUmsMemberMemberTagRelationModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&UmsMemberMemberTagRelation{}, id).Error
}

func (m *defaultUmsMemberMemberTagRelationModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsMemberMemberTagRelation, error) {
	var mm UmsMemberMemberTagRelation
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

func (m *defaultUmsMemberMemberTagRelationModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&UmsMemberMemberTagRelation{}).WithContext(ctx)
}

func (m *defaultUmsMemberMemberTagRelationModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultUmsMemberMemberTagRelationModel) FindPageListByBuilder(db *gorm.DB, page, pageSize int64) (*UmsMemberMemberTagRelationPagination, error) {
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(UmsMemberMemberTagRelationPagination)
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

	offset := int((page - 1) * pageSize)
	if err := db.Offset(offset).Limit(int(pageSize)).Find(&resp.Data).Error; err != nil {
		return nil, err
	}

	return resp, nil
}

func (m UmsMemberMemberTagRelationPagination) Pagination() UmsMemberMemberTagRelationPaginationBasic {
	return UmsMemberMemberTagRelationPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m UmsMemberMemberTagRelationPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultUmsMemberMemberTagRelationModel) Insert(ctx context.Context, data *UmsMemberMemberTagRelation) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultUmsMemberMemberTagRelationModel) Update(ctx context.Context, data *UmsMemberMemberTagRelation) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
