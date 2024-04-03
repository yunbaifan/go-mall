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
	pmsCommentModel interface {
		Insert(ctx context.Context, data *PmsComment) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsComment, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, page, pageSize int64) (*PmsCommentPagination, error)
		Update(ctx context.Context, data *PmsComment) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsCommentModel struct {
		table   string
		ormConn *gorm.DB
	}

	PmsComment struct {
		Id               int64          `json:"id" gorm:"primary_key"`
		ProductId        sql.NullInt64  `json:"product_id"`
		MemberNickName   sql.NullString `json:"member_nick_name"`
		ProductName      sql.NullString `json:"product_name"`
		Star             sql.NullInt64  `json:"star"`      // 评价星数：0->5
		MemberIp         sql.NullString `json:"member_ip"` // 评价的ip
		CreateTime       sql.NullTime   `json:"create_time"`
		ShowStatus       sql.NullInt64  `json:"show_status"`
		ProductAttribute sql.NullString `json:"product_attribute"` // 购买时的商品属性
		CollectCouont    sql.NullInt64  `json:"collect_couont"`
		ReadCount        sql.NullInt64  `json:"read_count"`
		Content          sql.NullString `json:"content"`
		Pics             sql.NullString `json:"pics"`        // 上传图片地址，以逗号隔开
		MemberIcon       sql.NullString `json:"member_icon"` // 评论用户头像
		ReplayCount      sql.NullInt64  `json:"replay_count"`
		PmsCommentRelation
	}

	PmsCommentPagination struct {
		Data []*PmsComment
		PmsCommentPaginationBasic
	}

	PmsCommentPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}
)

func newPmsCommentModel(db *gorm.DB) *defaultPmsCommentModel {
	return &defaultPmsCommentModel{
		table:   "`pms_comment`",
		ormConn: db,
	}
}

func (PmsComment) TableName() string {
	return "`pms_comment`"
}

func (m PmsComment) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultPmsCommentModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&PmsComment{}, id).Error
}

func (m *defaultPmsCommentModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsComment, error) {
	var mm PmsComment
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

func (m *defaultPmsCommentModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&PmsComment{}).WithContext(ctx)
}

func (m *defaultPmsCommentModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultPmsCommentModel) FindPageListByBuilder(db *gorm.DB, page, pageSize int64) (*PmsCommentPagination, error) {
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(PmsCommentPagination)
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

func (m PmsCommentPagination) Pagination() PmsCommentPaginationBasic {
	return PmsCommentPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m PmsCommentPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultPmsCommentModel) Insert(ctx context.Context, data *PmsComment) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultPmsCommentModel) Update(ctx context.Context, data *PmsComment) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
