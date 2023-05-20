package danmuauth

import (
	"net/http"

	"github.com/leaper-one/pkg/https/response"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/logic/danmuauth"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DanmuAuthStatusAddOneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddOneRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := danmuauth.NewDanmuAuthStatusAddOneLogic(r.Context(), svcCtx)
		resp, err := l.DanmuAuthStatusAddOne(&req)
		response.Response(w, resp, err)
	}
}
