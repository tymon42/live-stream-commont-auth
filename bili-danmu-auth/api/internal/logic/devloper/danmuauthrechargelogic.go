package devloper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/core"

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
	l.Logger.Infof("DanmuAuthRecharge,req: %v", req)

	var buid_string string = fmt.Sprintf("%v", l.ctx.Value("buid"))

	buid, err := strconv.Atoi(buid_string)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.BalanceDB.Charge(l.ctx, &core.Balance{Buid: buid}, req.Amount)
	if err != nil {
		return nil, err
	}

	return &types.RechargeResponse{Ok: true}, nil
}
