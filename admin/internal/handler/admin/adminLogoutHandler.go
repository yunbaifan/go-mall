package admin

import (
	"net/http"

	"github.com/yunbaifan/go-mall/admin/internal/logic/admin"
	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/lib/xcode"
)

func AdminLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewAdminLogoutLogic(r.Context(), svcCtx)
		err := l.AdminLogout()

		lang := svcCtx.Config.Lang
		var (
			formatResp interface{}
		)

		xcode.HttpResponse(r, w, formatResp, err, lang)
	}
}
