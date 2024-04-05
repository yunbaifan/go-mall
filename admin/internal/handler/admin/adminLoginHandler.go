package admin

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/yunbaifan/go-mall/admin/internal/logic/admin"
	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AdminLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		/*if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}*/
		if err := validator.New().StructCtx(r.Context(), &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := admin.NewAdminLoginLogic(r.Context(), svcCtx)
		resp, err := l.AdminLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
