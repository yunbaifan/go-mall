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
	umsMemberModel interface {
		Insert(ctx context.Context, data *UmsMember) error
		FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsMember, error)
		OrmSession(ctx context.Context) *gorm.DB
		Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
		FindPageListByBuilder(ctx context.Context, keyword *KeywordUmsMemberModel) (*UmsMemberPagination, error)
		FindOneByPhone(ctx context.Context, phone sql.NullString, preloadList ...string) (*UmsMember, error)
		FindOneByUsername(ctx context.Context, username sql.NullString, preloadList ...string) (*UmsMember, error)
		Update(ctx context.Context, data *UmsMember) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsMemberModel struct {
		table   string
		ormConn *gorm.DB
	}

	UmsMember struct {
		Id                    int64          `json:"id" gorm:"primary_key"`
		MemberLevelId         sql.NullInt64  `json:"member_level_id"`
		Username              sql.NullString `json:"username"`               // 用户名
		Password              sql.NullString `json:"password"`               // 密码
		Nickname              sql.NullString `json:"nickname"`               // 昵称
		Phone                 sql.NullString `json:"phone"`                  // 手机号码
		Status                sql.NullInt64  `json:"status"`                 // 帐号启用状态:0->禁用；1->启用
		CreateTime            sql.NullTime   `json:"create_time"`            // 注册时间
		Icon                  sql.NullString `json:"icon"`                   // 头像
		Gender                sql.NullInt64  `json:"gender"`                 // 性别：0->未知；1->男；2->女
		Birthday              sql.NullTime   `json:"birthday"`               // 生日
		City                  sql.NullString `json:"city"`                   // 所做城市
		Job                   sql.NullString `json:"job"`                    // 职业
		PersonalizedSignature sql.NullString `json:"personalized_signature"` // 个性签名
		SourceType            sql.NullInt64  `json:"source_type"`            // 用户来源
		Integration           sql.NullInt64  `json:"integration"`            // 积分
		Growth                sql.NullInt64  `json:"growth"`                 // 成长值
		LuckeyCount           sql.NullInt64  `json:"luckey_count"`           // 剩余抽奖次数
		HistoryIntegration    sql.NullInt64  `json:"history_integration"`    // 历史积分数量
		UmsMemberRelation
	}

	UmsMemberPagination struct {
		Data []*UmsMember
		UmsMemberPaginationBasic
	}

	UmsMemberPaginationBasic struct {
		CurrentPage int64
		PageSize    int64
		TotalCount  int64
		TotalPage   int64
	}

	KeywordUmsMemberModel struct {
		KeywordKey     string //like 关键字
		KeywordValue   string //like 值
		Page, PageSize int64
	}
)

func newUmsMemberModel(db *gorm.DB) *defaultUmsMemberModel {
	return &defaultUmsMemberModel{
		table:   "`ums_member`",
		ormConn: db,
	}
}

func (UmsMember) TableName() string {
	return "`ums_member`"
}

func (m UmsMember) ToProto(toValue interface{}) error {
	return copier.Copy(toValue, &m)
}

func (m *defaultUmsMemberModel) Delete(ctx context.Context, id int64) error {
	return m.OrmSession(ctx).Delete(&UmsMember{}, id).Error
}

func (m *defaultUmsMemberModel) FindOne(ctx context.Context, id int64, preloadList ...string) (*UmsMember, error) {
	var mm UmsMember
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

func (m *defaultUmsMemberModel) OrmSession(ctx context.Context) *gorm.DB {
	return m.ormConn.Model(&UmsMember{}).WithContext(ctx)
}

func (m *defaultUmsMemberModel) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return m.ormConn.WithContext(ctx).Transaction(fc, opts...)
}

func (m *defaultUmsMemberModel) FindPageListByBuilder(ctx context.Context, keyword *KeywordUmsMemberModel) (*UmsMemberPagination, error) {
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

	resp := new(UmsMemberPagination)
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

func (m UmsMemberPagination) Pagination() UmsMemberPaginationBasic {
	return UmsMemberPaginationBasic{
		CurrentPage: m.CurrentPage,
		PageSize:    m.PageSize,
		TotalPage:   m.TotalPage,
		TotalCount:  m.TotalCount,
	}
}

func (m UmsMemberPagination) ToProto(to any, toData any, toPagination any) error {
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

func (m *defaultUmsMemberModel) FindOneByPhone(ctx context.Context, phone sql.NullString, preloadList ...string) (*UmsMember, error) {
	var mm UmsMember
	db := m.OrmSession(ctx)
	for _, preload := range preloadList {
		db = db.Preload(preload)
	}
	err := db.First(&mm, "`phone` = ?", phone).Error
	switch err {
	case nil:
		return &mm, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

func (m *defaultUmsMemberModel) FindOneByUsername(ctx context.Context, username sql.NullString, preloadList ...string) (*UmsMember, error) {
	var mm UmsMember
	db := m.OrmSession(ctx)
	for _, preload := range preloadList {
		db = db.Preload(preload)
	}
	err := db.First(&mm, "`username` = ?", username).Error
	switch err {
	case nil:
		return &mm, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

func (m *defaultUmsMemberModel) Insert(ctx context.Context, data *UmsMember) error {
	return m.OrmSession(ctx).Create(data).Error
}

func (m *defaultUmsMemberModel) Update(ctx context.Context, newData *UmsMember) error {
	if newData.Id <= 0 {
		return gorm.ErrMissingWhereClause
	}
	return m.OrmSession(ctx).Where("id = ?", newData.Id).Save(newData).Error
}
