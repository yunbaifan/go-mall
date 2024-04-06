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
	umsGrowthChangeHistoryModel interface {
		Insert(ctx context.Context, data *UmsGrowthChangeHistory) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsGrowthChangeHistory, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, keyword *KeywordUmsGrowthChangeHistoryModel) (*UmsGrowthChangeHistoryPagination, error)
		Update(ctx context.Context, data *UmsGrowthChangeHistory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsGrowthChangeHistoryModel struct {
		table   string
		ormConn *gorm.DB
	}

	UmsGrowthChangeHistory struct {
		Id          int64          `json:"id" gorm:"primary_key"`
		MemberId    sql.NullInt64  `json:"member_id"`
		CreateTime  sql.NullTime   `json:"create_time"`
		ChangeType  sql.NullInt64  `json:"change_type"`  // 改变类型：0->增加；1->减少
		ChangeCount sql.NullInt64  `json:"change_count"` // 积分改变数量
		OperateMan  sql.NullString `json:"operate_man"`  // 操作人员
		OperateNote sql.NullString `json:"operate_note"` // 操作备注
		SourceType  sql.NullInt64  `json:"source_type"`  // 积分来源：0->购物；1->管理员修改
		UmsGrowthChangeHistoryRelation
	}

	UmsGrowthChangeHistoryPagination struct {
		Data []*UmsGrowthChangeHistory
		UmsGrowthChangeHistoryPaginationBasic
	}

	UmsGrowthChangeHistoryPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordUmsGrowthChangeHistoryModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newUmsGrowthChangeHistoryModel(db *gorm.DB) *defaultUmsGrowthChangeHistoryModel {
	return &defaultUmsGrowthChangeHistoryModel{
		table:   "`ums_growth_change_history`",
		ormConn: db,
	}
}

func (UmsGrowthChangeHistory) TableName() string {
	return "`ums_growth_change_history`"
}

func (m UmsGrowthChangeHistory) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultUmsGrowthChangeHistoryModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&UmsGrowthChangeHistory{}, id).Error
}

func (m *defaultUmsGrowthChangeHistoryModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsGrowthChangeHistory, error) {
	var mm UmsGrowthChangeHistory
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

func (m *defaultUmsGrowthChangeHistoryModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&UmsGrowthChangeHistory{}).WithContext(ctx)
}

func (m *defaultUmsGrowthChangeHistoryModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultUmsGrowthChangeHistoryModel) FindPageListByBuilder(db *gorm.DB, keyword *KeywordUmsGrowthChangeHistoryModel) (*UmsGrowthChangeHistoryPagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(UmsGrowthChangeHistoryPagination)
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

func (m UmsGrowthChangeHistoryPagination) Pagination() UmsGrowthChangeHistoryPaginationBasic {
	return UmsGrowthChangeHistoryPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m UmsGrowthChangeHistoryPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultUmsGrowthChangeHistoryModel) Insert(ctx context.Context, data *UmsGrowthChangeHistory) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultUmsGrowthChangeHistoryModel) Update(ctx context.Context, data *UmsGrowthChangeHistory) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
