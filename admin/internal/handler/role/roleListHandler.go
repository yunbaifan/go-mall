package role

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/yunbaifan/go-mall/admin/internal/logic/role"
	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"
	"github.com/yunbaifan/go-mall/lib/xcode"
)

func RoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if err := validator.New().StructCtx(r.Context(), &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := role.NewRoleListLogic(r.Context(), svcCtx)
		resp, err := l.RoleList(&req)

		lang := svcCtx.Config.Lang
		var (
			formatResp interface{}
		)
		formatResp = xcode.SuccessResponse(resp, lang)

		xcode.HttpResponse(r, w, formatResp, err, lang)
	}
}
