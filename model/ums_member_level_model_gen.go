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
	umsMemberLevelModel interface {
		Insert(ctx context.Context, data *UmsMemberLevel) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsMemberLevel, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ormSession *gorm.DB, page, pageSize int64) (*UmsMemberLevelPagination, error)
		Update(ctx context.Context, data *UmsMemberLevel) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsMemberLevelModel struct {
		table   string
		ormConn *gorm.DB
	}

	UmsMemberLevel struct {
		Id                    int64           `json:"id" gorm:"primary_key"`
		Name                  sql.NullString  `json:"name"`
		GrowthPoint           sql.NullInt64   `json:"growth_point"`
		DefaultStatus         sql.NullInt64   `json:"default_status"`          // 是否为默认等级：0->不是；1->是
		FreeFreightPoint      sql.NullFloat64 `json:"free_freight_point"`      // 免运费标准
		CommentGrowthPoint    sql.NullInt64   `json:"comment_growth_point"`    // 每次评价获取的成长值
		PriviledgeFreeFreight sql.NullInt64   `json:"priviledge_free_freight"` // 是否有免邮特权
		PriviledgeSignIn      sql.NullInt64   `json:"priviledge_sign_in"`      // 是否有签到特权
		PriviledgeComment     sql.NullInt64   `json:"priviledge_comment"`      // 是否有评论获奖励特权
		PriviledgePromotion   sql.NullInt64   `json:"priviledge_promotion"`    // 是否有专享活动特权
		PriviledgeMemberPrice sql.NullInt64   `json:"priviledge_member_price"` // 是否有会员价格特权
		PriviledgeBirthday    sql.NullInt64   `json:"priviledge_birthday"`     // 是否有生日特权
		Note                  sql.NullString  `json:"note"`
		UmsMemberLevelRelation
	}

	UmsMemberLevelPagination struct {
		Data []*UmsMemberLevel
		UmsMemberLevelPaginationBasic
	}

	UmsMemberLevelPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}
)

func newUmsMemberLevelModel(db *gorm.DB) *defaultUmsMemberLevelModel {
	return &defaultUmsMemberLevelModel{
		table:   "`ums_member_level`",
		ormConn: db,
	}
}

func (UmsMemberLevel) TableName() string {
	return "`ums_member_level`"
}

func (m UmsMemberLevel) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultUmsMemberLevelModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&UmsMemberLevel{}, id).Error
}

func (m *defaultUmsMemberLevelModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsMemberLevel, error) {
	var mm UmsMemberLevel
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

func (m *defaultUmsMemberLevelModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&UmsMemberLevel{}).WithContext(ctx)
}

func (m *defaultUmsMemberLevelModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultUmsMemberLevelModel) FindPageListByBuilder(db *gorm.DB, page, pageSize int64) (*UmsMemberLevelPagination, error) {
	// 总行数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 总页数
	totalPage := int64(math.Ceil(float64(totalCount) / float64(pageSize)))

	resp := new(UmsMemberLevelPagination)
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

func (m UmsMemberLevelPagination) Pagination() UmsMemberLevelPaginationBasic {
	return UmsMemberLevelPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m UmsMemberLevelPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultUmsMemberLevelModel) Insert(ctx context.Context, data *UmsMemberLevel) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultUmsMemberLevelModel) Update(ctx context.Context, data *UmsMemberLevel) error {
	if data.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", data.Id).Save(data).Error
}