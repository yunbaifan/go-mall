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
	pmsCommentReplayModel interface {
		Insert(ctx context.Context, data *PmsCommentReplay) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsCommentReplay, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, keyword *KeywordPmsCommentReplayModel) (*PmsCommentReplayPagination, error)
		Update(ctx context.Context, data *PmsCommentReplay) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsCommentReplayModel struct {
		table   string
		ormConn *gorm.DB
	}

	PmsCommentReplay struct {
		Id             int64          `json:"id" gorm:"primary_key"`
		CommentId      sql.NullInt64  `json:"comment_id"`
		MemberNickName sql.NullString `json:"member_nick_name"`
		MemberIcon     sql.NullString `json:"member_icon"`
		Content        sql.NullString `json:"content"`
		CreateTime     sql.NullTime   `json:"create_time"`
		Type           sql.NullInt64  `json:"type"` // 评论人员类型；0->会员；1->管理员
		PmsCommentReplayRelation
	}

	PmsCommentReplayPagination struct {
		Data []*PmsCommentReplay
		PmsCommentReplayPaginationBasic
	}

	PmsCommentReplayPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordPmsCommentReplayModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newPmsCommentReplayModel(db *gorm.DB) *defaultPmsCommentReplayModel {
	return &defaultPmsCommentReplayModel{
		table:   "`pms_comment_replay`",
		ormConn: db,
	}
}

func (PmsCommentReplay) TableName() string {
	return "`pms_comment_replay`"
}

func (m PmsCommentReplay) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultPmsCommentReplayModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&PmsCommentReplay{}, id).Error
}

func (m *defaultPmsCommentReplayModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsCommentReplay, error) {
	var mm PmsCommentReplay
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

func (m *defaultPmsCommentReplayModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&PmsCommentReplay{}).WithContext(ctx)
}

func (m *defaultPmsCommentReplayModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultPmsCommentReplayModel) FindPageListByBuilder(db *gorm.DB, keyword *KeywordPmsCommentReplayModel) (*PmsCommentReplayPagination, error) {
	page := keyword.Page
	pageSize := keyword.PageSize
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(PmsCommentReplayPagination)
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

func (m PmsCommentReplayPagination) Pagination() PmsCommentReplayPaginationBasic {
	return PmsCommentReplayPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m PmsCommentReplayPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultPmsCommentReplayModel) Insert(ctx context.Context, data *PmsCommentReplay) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultPmsCommentReplayModel) Update(ctx context.Context, data *PmsCommentReplay) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}
