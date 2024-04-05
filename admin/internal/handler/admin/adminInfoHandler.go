package admin

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/yunbaifan/go-mall/admin/internal/logic/admin"
	"github.com/yunbaifan/go-mall/admin/internal/svc"
)

func AdminInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewAdminInfoLogic(r.Context(), svcCtx)
		resp, err := l.AdminInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
