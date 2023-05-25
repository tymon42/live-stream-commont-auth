package danmuauth

import (
	"context"
	"errors"

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
	l.Logger.Infof("DanmuAUthStatusAddOne,req: %v", req)

	if req.ApiKey != l.svcCtx.Config.Worker.ApiKey {
		return nil, errors.New("worker api_key error")
	}

	da, err := l.svcCtx.DanmuAuthDB.FindByBuidVCode(l.ctx, req.Buid, req.Vcode, l.svcCtx.Config.DanmuAuth.VCodeExpire)
	if err != nil {
		return nil, err
	} else if da == nil && err == nil {
		return nil, errors.New("vcode not found")
	}

	// add one
	err = l.svcCtx.DanmuAuthDB.AddVerifiedCount(l.ctx, da)
	if err != nil {
		return nil, err
	}

	return &types.AddOneResponse{Status: da.VerifiedCount}, nil
}
