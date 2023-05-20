package danmuauth

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthStatusAddOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthStatusAddOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthStatusAddOneLogic {
	return &DanmuAuthStatusAddOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthStatusAddOneLogic) DanmuAuthStatusAddOne(req *types.AddOneRequest) (resp *types.AddOneResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
