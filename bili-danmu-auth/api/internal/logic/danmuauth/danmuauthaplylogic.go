package danmuauth

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthAplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthAplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthAplyLogic {
	return &DanmuAuthAplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthAplyLogic) DanmuAuthAply(req *types.ApplyRequest) (resp *types.ApplyResponse, err error) {
	

	return
}
