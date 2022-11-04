package danmuauth

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthStatusLogic {
	return &DanmuAuthStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthStatusLogic) DanmuAuthStatus(req *types.StatusRequest) (resp *types.StatusResponse, err error) {
	l.Logger.Info("DanmuAuthStatus", req)

	danmuAuth, err := l.svcCtx.DanmuAuthDB.FindByUUIDBuidVCode(l.ctx, req.Client_uuid, uint(req.Buid), req.Vcode)
	if err != nil {
		return nil, err
	}
	if danmuAuth == nil && err == nil {
		return &types.StatusResponse{
			Verify_count: -1,
		}, nil
	}
	if danmuAuth.VerifiedCount >= int(req.Count) {
		err = l.svcCtx.DanmuAuthDB.Delete(l.ctx, danmuAuth.ID)
		if err != nil {
			return nil, err
		}
	}
	return &types.StatusResponse{
		Verify_count: int64(danmuAuth.VerifiedCount),
	}, nil
}
