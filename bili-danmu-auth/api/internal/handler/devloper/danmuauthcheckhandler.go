package devloper

import (
	"net/http"

	"github.com/leaper-one/pkg/https/response"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/logic/devloper"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DanmuAuthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := devloper.NewDanmuAuthCheckLogic(r.Context(), svcCtx)
		resp, err := l.DanmuAuthCheck(&req)
		response.Response(w, resp, err)
	}
}
