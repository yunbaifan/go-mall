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
	pmsFeightTemplateModel interface {
		Insert(ctx context.Context, data *PmsFeightTemplate) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsFeightTemplate, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ctx context.Context, keyword *KeywordPmsFeightTemplateModel) (*PmsFeightTemplatePagination, error)
		Update(ctx context.Context, data *PmsFeightTemplate) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsFeightTemplateModel struct {
		table   string
		ormConn *gorm.DB
	}

	PmsFeightTemplate struct {
		Id             int64           `json:"id" gorm:"primary_key"`
		Name           sql.NullString  `json:"name"`
		ChargeType     sql.NullInt64   `json:"charge_type"`  // 计费类型:0->按重量；1->按件数
		FirstWeight    sql.NullFloat64 `json:"first_weight"` // 首重kg
		FirstFee       sql.NullFloat64 `json:"first_fee"`    // 首费（元）
		ContinueWeight sql.NullFloat64 `json:"continue_weight"`
		ContinmeFee    sql.NullFloat64 `json:"continme_fee"`
		Dest           sql.NullString  `json:"dest"` // 目的地（省、市）
		PmsFeightTemplateRelation
	}

	PmsFeightTemplatePagination struct {
		Data []*PmsFeightTemplate
		PmsFeightTemplatePaginationBasic
	}

	PmsFeightTemplatePaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordPmsFeightTemplateModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newPmsFeightTemplateModel(db *gorm.DB) *defaultPmsFeightTemplateModel {
	return &defaultPmsFeightTemplateModel{
		table:   "`pms_feight_template`",
		ormConn: db,
	}
}

func (PmsFeightTemplate) TableName() string {
	return "`pms_feight_template`"
}

func (m PmsFeightTemplate) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultPmsFeightTemplateModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&PmsFeightTemplate{}, id).Error
}

func (m *defaultPmsFeightTemplateModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsFeightTemplate, error) {
	var mm PmsFeightTemplate
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

func (m *defaultPmsFeightTemplateModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&PmsFeightTemplate{}).WithContext(ctx)
}

func (m *defaultPmsFeightTemplateModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultPmsFeightTemplateModel) FindPageListByBuilder(ctx context.Context, keyword *KeywordPmsFeightTemplateModel) (*PmsFeightTemplatePagination, error) {
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

	resp := new(PmsFeightTemplatePagination)
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

func (m PmsFeightTemplatePagination) Pagination() PmsFeightTemplatePaginationBasic {
	return PmsFeightTemplatePaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m PmsFeightTemplatePagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultPmsFeightTemplateModel) Insert(ctx context.Context, data *PmsFeightTemplate) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultPmsFeightTemplateModel) Update(ctx context.Context, data *PmsFeightTemplate) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
