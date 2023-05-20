package devloper

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthGetKeyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthGetKeyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthGetKeyListLogic {
	return &DanmuAuthGetKeyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthGetKeyListLogic) DanmuAuthGetKeyList(req *types.GetKeyListRequest) (resp *types.GetKeyListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
