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
	pmsBrandModel interface {
		Insert(ctx context.Context, data *PmsBrand) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsBrand, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, keyword *KeywordPmsBrandModel) (*PmsBrandPagination, error)
		Update(ctx context.Context, data *PmsBrand) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsBrandModel struct {
		table   string
		ormConn *gorm.DB
	}

	PmsBrand struct {
		Id                  int64          `json:"id" gorm:"primary_key"`
		Name                sql.NullString `json:"name"`
		FirstLetter         sql.NullString `json:"first_letter"` // 首字母
		Sort                sql.NullInt64  `json:"sort"`
		FactoryStatus       sql.NullInt64  `json:"factory_status"` // 是否为品牌制造商：0->不是；1->是
		ShowStatus          sql.NullInt64  `json:"show_status"`
		ProductCount        sql.NullInt64  `json:"product_count"`         // 产品数量
		ProductCommentCount sql.NullInt64  `json:"product_comment_count"` // 产品评论数量
		Logo                sql.NullString `json:"logo"`                  // 品牌logo
		BigPic              sql.NullString `json:"big_pic"`               // 专区大图
		BrandStory          sql.NullString `json:"brand_story"`           // 品牌故事
		PmsBrandRelation
	}

	PmsBrandPagination struct {
		Data []*PmsBrand
		PmsBrandPaginationBasic
	}

	PmsBrandPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordPmsBrandModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newPmsBrandModel(db *gorm.DB) *defaultPmsBrandModel {
	return &defaultPmsBrandModel{
		table:   "`pms_brand`",
		ormConn: db,
	}
}

func (PmsBrand) TableName() string {
	return "`pms_brand`"
}

func (m PmsBrand) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultPmsBrandModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&PmsBrand{}, id).Error
}

func (m *defaultPmsBrandModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsBrand, error) {
	var mm PmsBrand
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

func (m *defaultPmsBrandModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&PmsBrand{}).WithContext(ctx)
}

func (m *defaultPmsBrandModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultPmsBrandModel) FindPageListByBuilder(db *gorm.DB, keyword *KeywordPmsBrandModel) (*PmsBrandPagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(PmsBrandPagination)
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

func (m PmsBrandPagination) Pagination() PmsBrandPaginationBasic {
	return PmsBrandPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m PmsBrandPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultPmsBrandModel) Insert(ctx context.Context, data *PmsBrand) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultPmsBrandModel) Update(ctx context.Context, data *PmsBrand) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
