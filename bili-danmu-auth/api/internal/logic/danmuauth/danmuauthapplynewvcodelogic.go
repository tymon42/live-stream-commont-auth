package danmuauth

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthApplyNewVCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthApplyNewVCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthApplyNewVCodeLogic {
	return &DanmuAuthApplyNewVCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthApplyNewVCodeLogic) DanmuAuthApplyNewVCode(req *types.ApplyNewVCodeRequest) (resp *types.ApplyNewVCodeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
