package devloper

import (
	"context"

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

	blc, err := l.svcCtx.BalanceDB.FindByBuid(l.ctx, req.Buid)
	if err != nil {
		return nil, err
	} else if blc == nil && err == nil {
		err = l.svcCtx.BalanceDB.Save(l.ctx, &core.Balance{Buid: req.Buid, Balance: 50})
		if err != nil {
			return nil, err
		}
	}

	err = l.svcCtx.BalanceDB.Charge(l.ctx, blc, req.Amount)
	if err != nil {
		return nil, err
	}
	l.Logger.Infof("DanmuAuthRecharge, charge success")

	return &types.RechargeResponse{Ok: true}, nil
}
