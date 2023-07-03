package danmuauth

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthVCodeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthVCodeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthVCodeInfoLogic {
	return &DanmuAuthVCodeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthVCodeInfoLogic) DanmuAuthVCodeInfo(req *types.VCodeInfoRequest) (resp *types.VCodeInfoResponse, err error) {
	l.Logger.Infof("DanmuAuthVCodeInfo: %v", req)

	danmuAuth, err := l.svcCtx.DanmuAuthDB.FindByClientIDAndBuid(l.ctx, req.ClientID, req.Buid, int(l.svcCtx.Config.DanmuAuth.VCodeExpire))
	if err != nil {
		l.Logger.Errorf("find danmu auth failed, err: %v", err)
		return nil, err
	}

	resp = &types.VCodeInfoResponse{
		Count:    danmuAuth.VerifiedCount,
		ClientID: danmuAuth.ClientID,
	}

	l.Logger.Infof("DanmuAuthVCodeInfo success, resp: %v", resp)

	return resp, nil
}
