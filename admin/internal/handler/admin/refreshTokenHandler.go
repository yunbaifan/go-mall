package admin

import (
	"net/http"

	"github.com/yunbaifan/go-mall/admin/internal/logic/admin"
	"github.com/yunbaifan/go-mall/admin/internal/svc"
	"github.com/yunbaifan/go-mall/lib/xcode"
)

func RefreshTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewRefreshTokenLogic(r.Context(), svcCtx)
		resp, err := l.RefreshToken()

		lang := svcCtx.Config.Lang
		var (
			formatResp interface{}
		)
		formatResp = xcode.SuccessResponse(resp, lang)

		xcode.HttpResponse(r, w, formatResp, err, lang)
	}
}
