package devloper

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthCheckLogic {
	return &DanmuAuthCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthCheckLogic) DanmuAuthCheck(req *types.CheckRequest) (resp *types.CheckResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
