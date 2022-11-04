package danmuauth

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthStatusLogic {
	return &DanmuAuthStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthStatusLogic) DanmuAuthStatus(req *types.StatusRequest) (resp *types.StatusResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
