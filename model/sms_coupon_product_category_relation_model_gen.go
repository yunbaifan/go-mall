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
	smsCouponProductCategoryRelationModel interface {
		Insert(ctx context.Context, data *SmsCouponProductCategoryRelation) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*SmsCouponProductCategoryRelation, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, page, pageSize int64) (*SmsCouponProductCategoryRelationPagination, error)
		Update(ctx context.Context, data *SmsCouponProductCategoryRelation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSmsCouponProductCategoryRelationModel struct {
		table   string
		ormConn *gorm.DB
	}

	SmsCouponProductCategoryRelation struct {
		Id                  int64          `json:"id" gorm:"primary_key"`
		CouponId            sql.NullInt64  `json:"coupon_id"`
		ProductCategoryId   sql.NullInt64  `json:"product_category_id"`
		ProductCategoryName sql.NullString `json:"product_category_name"` // 产品分类名称
		ParentCategoryName  sql.NullString `json:"parent_category_name"`  // 父分类名称
		SmsCouponProductCategoryRelationRelation
	}

	SmsCouponProductCategoryRelationPagination struct {
		Data []*SmsCouponProductCategoryRelation
		SmsCouponProductCategoryRelationPaginationBasic
	}

	SmsCouponProductCategoryRelationPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}
)

func newSmsCouponProductCategoryRelationModel(db *gorm.DB) *defaultSmsCouponProductCategoryRelationModel {
	return &defaultSmsCouponProductCategoryRelationModel{
		table:   "`sms_coupon_product_category_relation`",
		ormConn: db,
	}
}

func (SmsCouponProductCategoryRelation) TableName() string {
	return "`sms_coupon_product_category_relation`"
}

func (m SmsCouponProductCategoryRelation) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultSmsCouponProductCategoryRelationModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&SmsCouponProductCategoryRelation{}, id).Error
}

func (m *defaultSmsCouponProductCategoryRelationModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*SmsCouponProductCategoryRelation, error) {
	var mm SmsCouponProductCategoryRelation
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

func (m *defaultSmsCouponProductCategoryRelationModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&SmsCouponProductCategoryRelation{}).WithContext(ctx)
}

func (m *defaultSmsCouponProductCategoryRelationModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultSmsCouponProductCategoryRelationModel) FindPageListByBuilder(db *gorm.DB, page, pageSize int64) (*SmsCouponProductCategoryRelationPagination, error) {
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(SmsCouponProductCategoryRelationPagination)
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

func (m SmsCouponProductCategoryRelationPagination) Pagination() SmsCouponProductCategoryRelationPaginationBasic {
	return SmsCouponProductCategoryRelationPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m SmsCouponProductCategoryRelationPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultSmsCouponProductCategoryRelationModel) Insert(ctx context.Context, data *SmsCouponProductCategoryRelation) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultSmsCouponProductCategoryRelationModel) Update(ctx context.Context, data *SmsCouponProductCategoryRelation) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}