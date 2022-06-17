package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"temp-service/apps/app/api/internal/logic/user"
	"temp-service/apps/app/api/internal/svc"
	"temp-service/apps/app/api/internal/types"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
