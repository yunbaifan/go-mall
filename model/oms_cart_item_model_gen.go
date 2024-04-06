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
	omsCartItemModel interface {
		Insert(ctx context.Context, data *OmsCartItem) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*OmsCartItem, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ctx context.Context, keyword *KeywordOmsCartItemModel) (*OmsCartItemPagination, error)
		Update(ctx context.Context, data *OmsCartItem) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOmsCartItemModel struct {
		table   string
		ormConn *gorm.DB
	}

	OmsCartItem struct {
		Id                int64           `json:"id" gorm:"primary_key"`
		ProductId         sql.NullInt64   `json:"product_id"`
		ProductSkuId      sql.NullInt64   `json:"product_sku_id"`
		MemberId          sql.NullInt64   `json:"member_id"`
		Quantity          sql.NullInt64   `json:"quantity"`            // 购买数量
		Price             sql.NullFloat64 `json:"price"`               // 添加到购物车的价格
		ProductPic        sql.NullString  `json:"product_pic"`         // 商品主图
		ProductName       sql.NullString  `json:"product_name"`        // 商品名称
		ProductSubTitle   sql.NullString  `json:"product_sub_title"`   // 商品副标题（卖点）
		ProductSkuCode    sql.NullString  `json:"product_sku_code"`    // 商品sku条码
		MemberNickname    sql.NullString  `json:"member_nickname"`     // 会员昵称
		CreateDate        sql.NullTime    `json:"create_date"`         // 创建时间
		ModifyDate        sql.NullTime    `json:"modify_date"`         // 修改时间
		DeleteStatus      int64           `json:"delete_status"`       // 是否删除
		ProductCategoryId sql.NullInt64   `json:"product_category_id"` // 商品分类
		ProductBrand      sql.NullString  `json:"product_brand"`
		ProductSn         sql.NullString  `json:"product_sn"`
		ProductAttr       sql.NullString  `json:"product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
		OmsCartItemRelation
	}

	OmsCartItemPagination struct {
		Data []*OmsCartItem
		OmsCartItemPaginationBasic
	}

	OmsCartItemPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordOmsCartItemModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newOmsCartItemModel(db *gorm.DB) *defaultOmsCartItemModel {
	return &defaultOmsCartItemModel{
		table:   "`oms_cart_item`",
		ormConn: db,
	}
}

func (OmsCartItem) TableName() string {
	return "`oms_cart_item`"
}

func (m OmsCartItem) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultOmsCartItemModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&OmsCartItem{}, id).Error
}

func (m *defaultOmsCartItemModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*OmsCartItem, error) {
	var mm OmsCartItem
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

func (m *defaultOmsCartItemModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&OmsCartItem{}).WithContext(ctx)
}

func (m *defaultOmsCartItemModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultOmsCartItemModel) FindPageListByBuilder(ctx context.Context, keyword *KeywordOmsCartItemModel) (*OmsCartItemPagination, error) {
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

	resp := new(OmsCartItemPagination)
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

func (m OmsCartItemPagination) Pagination() OmsCartItemPaginationBasic {
	return OmsCartItemPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m OmsCartItemPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultOmsCartItemModel) Insert(ctx context.Context, data *OmsCartItem) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultOmsCartItemModel) Update(ctx context.Context, data *OmsCartItem) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
