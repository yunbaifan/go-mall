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
	pmsAlbumPicModel interface {
		Insert(ctx context.Context, data *PmsAlbumPic) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsAlbumPic, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, page, pageSize int64) (*PmsAlbumPicPagination, error)
		Update(ctx context.Context, data *PmsAlbumPic) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsAlbumPicModel struct {
		table   string
		ormConn *gorm.DB
	}

	PmsAlbumPic struct {
		Id      int64          `json:"id" gorm:"primary_key"`
		AlbumId sql.NullInt64  `json:"album_id"`
		Pic     sql.NullString `json:"pic"`
		PmsAlbumPicRelation
	}

	PmsAlbumPicPagination struct {
		Data []*PmsAlbumPic
		PmsAlbumPicPaginationBasic
	}

	PmsAlbumPicPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}
)

func newPmsAlbumPicModel(db *gorm.DB) *defaultPmsAlbumPicModel {
	return &defaultPmsAlbumPicModel{
		table:   "`pms_album_pic`",
		ormConn: db,
	}
}

func (PmsAlbumPic) TableName() string {
	return "`pms_album_pic`"
}

func (m PmsAlbumPic) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultPmsAlbumPicModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&PmsAlbumPic{}, id).Error
}

func (m *defaultPmsAlbumPicModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*PmsAlbumPic, error) {
	var mm PmsAlbumPic
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

func (m *defaultPmsAlbumPicModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&PmsAlbumPic{}).WithContext(ctx)
}

func (m *defaultPmsAlbumPicModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultPmsAlbumPicModel) FindPageListByBuilder(db *gorm.DB, page, pageSize int64) (*PmsAlbumPicPagination, error) {
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(PmsAlbumPicPagination)
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

func (m PmsAlbumPicPagination) Pagination() PmsAlbumPicPaginationBasic {
	return PmsAlbumPicPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m PmsAlbumPicPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultPmsAlbumPicModel) Insert(ctx context.Context, data *PmsAlbumPic) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultPmsAlbumPicModel) Update(ctx context.Context, data *PmsAlbumPic) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}