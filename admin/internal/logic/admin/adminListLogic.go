package admin

import (
	"context"
	"github.com/yunbaifan/go-mall/model"

	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminListLogic {
	return &AdminListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminListLogic) AdminList(req *types.AdminListRequest) (resp *types.AdminListResponse, err error) {
	res, err := l.svcCtx.UmsAdminModel.FindPageListByBuilder(l.ctx, &model.KeywordUmsAdminModel{
		KeywordKey:   "username",
		KeywordValue: req.Keyword,
		Page:         req.PageNum,
		PageSize:     req.PageSize,
	})
	if err != nil {
		l.Logger.Errorf("AdminList.UmsAdminModel.FindPageListByBuilder error: %v req:%v", err, req)
		return
	}
	list := make([]types.AdminResponse, 0, len(res.Data))
	for k := range res.Data {
		val := res.Data[k]
		list = append(list, types.AdminResponse{
			ID:         val.Id,
			Username:   val.Username.String,
			CreateTime: val.CreateTime.Time.Format("2006-01-02 15:04:05"),
			Status:     val.Status,
			LoginTime:  val.LoginTime.Time.Format("2006-01-02 15:04:05"),
			NickName:   val.NickName.String,
			Icon:       val.Icon.String,
			Email:      val.Email.String,
			Password:   val.Password.String,
			Note:       val.Note.String,
		})
	}
	return &types.AdminListResponse{
		Total:     res.UmsAdminPaginationBasic.TotalCount,
		List:      list,
		PageNum:   res.UmsAdminPaginationBasic.CurrentPage,
		PageSize:  res.UmsAdminPaginationBasic.PageSize,
		TotalPage: res.UmsAdminPaginationBasic.TotalPage,
	}, nil

}
