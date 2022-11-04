package danmuauth

import (
	"context"
	"errors"

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
	l.Logger.Info("DanmuAuthSubmit", req)

	danmuAuth, err := l.svcCtx.DanmuAuthDB.FindByBuidVCode(l.ctx, uint(req.Buid), req.Vcode)
	if err != nil {
		return nil, err
	}
	if danmuAuth == nil && err == nil {
		return &types.SubmitResponse{}, errors.New("no danmuAuth found")
	}

	danmuAuth.VerifiedCount++
	err = l.svcCtx.DanmuAuthDB.Save(l.ctx, danmuAuth)
	if err != nil {
		return nil, err
	}
	return
}
