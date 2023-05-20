package devloper

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthDelKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthDelKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthDelKeyLogic {
	return &DanmuAuthDelKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthDelKeyLogic) DanmuAuthDelKey(req *types.DeleteKeyRequest) (resp *types.DeleteKeyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
