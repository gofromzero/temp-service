package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"temp-service/apps/app/api/internal/logic"
	"temp-service/apps/app/api/internal/svc"
	"temp-service/apps/app/api/internal/types"
	"temp-service/pkg/result"
)

func ApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewApiLogic(r.Context(), svcCtx)
		resp, err := l.Api(&req)

		result.Response(w, resp, err)
	}
}
