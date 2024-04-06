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
	omsOrderItemModel interface {
		Insert(ctx context.Context, data *OmsOrderItem) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*OmsOrderItem, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, keyword *KeywordOmsOrderItemModel) (*OmsOrderItemPagination, error)
		Update(ctx context.Context, data *OmsOrderItem) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOmsOrderItemModel struct {
		table   string
		ormConn *gorm.DB
	}

	OmsOrderItem struct {
		Id                int64           `json:"id" gorm:"primary_key"`
		OrderId           sql.NullInt64   `json:"order_id"` // 订单id
		OrderSn           sql.NullString  `json:"order_sn"` // 订单编号
		ProductId         sql.NullInt64   `json:"product_id"`
		ProductPic        sql.NullString  `json:"product_pic"`
		ProductName       sql.NullString  `json:"product_name"`
		ProductBrand      sql.NullString  `json:"product_brand"`
		ProductSn         sql.NullString  `json:"product_sn"`
		ProductPrice      sql.NullFloat64 `json:"product_price"`       // 销售价格
		ProductQuantity   sql.NullInt64   `json:"product_quantity"`    // 购买数量
		ProductSkuId      sql.NullInt64   `json:"product_sku_id"`      // 商品sku编号
		ProductSkuCode    sql.NullString  `json:"product_sku_code"`    // 商品sku条码
		ProductCategoryId sql.NullInt64   `json:"product_category_id"` // 商品分类id
		PromotionName     sql.NullString  `json:"promotion_name"`      // 商品促销名称
		PromotionAmount   sql.NullFloat64 `json:"promotion_amount"`    // 商品促销分解金额
		CouponAmount      sql.NullFloat64 `json:"coupon_amount"`       // 优惠券优惠分解金额
		IntegrationAmount sql.NullFloat64 `json:"integration_amount"`  // 积分优惠分解金额
		RealAmount        sql.NullFloat64 `json:"real_amount"`         // 该商品经过优惠后的分解金额
		GiftIntegration   int64           `json:"gift_integration"`
		GiftGrowth        int64           `json:"gift_growth"`
		ProductAttr       sql.NullString  `json:"product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
		OmsOrderItemRelation
	}

	OmsOrderItemPagination struct {
		Data []*OmsOrderItem
		OmsOrderItemPaginationBasic
	}

	OmsOrderItemPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordOmsOrderItemModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newOmsOrderItemModel(db *gorm.DB) *defaultOmsOrderItemModel {
	return &defaultOmsOrderItemModel{
		table:   "`oms_order_item`",
		ormConn: db,
	}
}

func (OmsOrderItem) TableName() string {
	return "`oms_order_item`"
}

func (m OmsOrderItem) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultOmsOrderItemModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&OmsOrderItem{}, id).Error
}

func (m *defaultOmsOrderItemModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*OmsOrderItem, error) {
	var mm OmsOrderItem
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

func (m *defaultOmsOrderItemModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&OmsOrderItem{}).WithContext(ctx)
}

func (m *defaultOmsOrderItemModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultOmsOrderItemModel) FindPageListByBuilder(db *gorm.DB, keyword *KeywordOmsOrderItemModel) (*OmsOrderItemPagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(OmsOrderItemPagination)
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

func (m OmsOrderItemPagination) Pagination() OmsOrderItemPaginationBasic {
	return OmsOrderItemPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m OmsOrderItemPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultOmsOrderItemModel) Insert(ctx context.Context, data *OmsOrderItem) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultOmsOrderItemModel) Update(ctx context.Context, data *OmsOrderItem) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
