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
	pmsSkuStockModel interface {
		Insert(ctx context.Context, data *PmsSkuStock) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsSkuStock, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, page, pageSize int64) (*PmsSkuStockPagination, error)
		Update(ctx context.Context, data *PmsSkuStock) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsSkuStockModel struct {
		table   string
		ormConn *gorm.DB
	}

	PmsSkuStock struct {
		Id             int64           `json:"id" gorm:"primary_key"`
		ProductId      sql.NullInt64   `json:"product_id"`
		SkuCode        string          `json:"sku_code"` // sku编码
		Price          sql.NullFloat64 `json:"price"`
		Stock          int64           `json:"stock"`           // 库存
		LowStock       sql.NullInt64   `json:"low_stock"`       // 预警库存
		Pic            sql.NullString  `json:"pic"`             // 展示图片
		Sale           sql.NullInt64   `json:"sale"`            // 销量
		PromotionPrice sql.NullFloat64 `json:"promotion_price"` // 单品促销价格
		LockStock      int64           `json:"lock_stock"`      // 锁定库存
		SpData         sql.NullString  `json:"sp_data"`         // 商品销售属性，json格式
		PmsSkuStockRelation
	}

	PmsSkuStockPagination struct {
		Data []*PmsSkuStock
		PmsSkuStockPaginationBasic
	}

	PmsSkuStockPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}
)

func newPmsSkuStockModel(db *gorm.DB) *defaultPmsSkuStockModel {
	return &defaultPmsSkuStockModel{
		table:   "`pms_sku_stock`",
		ormConn: db,
	}
}

func (PmsSkuStock) TableName() string {
	return "`pms_sku_stock`"
}

func (m PmsSkuStock) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultPmsSkuStockModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&PmsSkuStock{}, id).Error
}

func (m *defaultPmsSkuStockModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsSkuStock, error) {
	var mm PmsSkuStock
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

func (m *defaultPmsSkuStockModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&PmsSkuStock{}).WithContext(ctx)
}

func (m *defaultPmsSkuStockModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultPmsSkuStockModel) FindPageListByBuilder(db *gorm.DB, page, pageSize int64) (*PmsSkuStockPagination, error) {
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(PmsSkuStockPagination)
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

func (m PmsSkuStockPagination) Pagination() PmsSkuStockPaginationBasic {
	return PmsSkuStockPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m PmsSkuStockPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultPmsSkuStockModel) Insert(ctx context.Context, data *PmsSkuStock) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultPmsSkuStockModel) Update(ctx context.Context, data *PmsSkuStock) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
