package danmuauth

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthSubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthSubmitLogic {
	return &DanmuAuthSubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthSubmitLogic) DanmuAuthSubmit(req *types.SubmitRequest) (resp *types.SubmitResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
