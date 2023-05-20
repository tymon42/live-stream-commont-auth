package devloper

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthRechargeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthRechargeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthRechargeLogic {
	return &DanmuAuthRechargeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthRechargeLogic) DanmuAuthRecharge(req *types.RechargeRequest) (resp *types.RechargeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
