package danmuauth

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"
	"github.com/tymon42/live-stream-commont-auth/vcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthAplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthAplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthAplyLogic {
	return &DanmuAuthAplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthAplyLogic) DanmuAuthAply(req *types.ApplyRequest) (resp *types.ApplyResponse, err error) {
	l.Logger.Info("DanmuAuthAply", req)

	danmuAuth, err := l.svcCtx.DanmuAuthDB.FindByUUIDBuid(l.ctx, req.Client_uuid, uint(req.Buid))
	if err != nil {
		return nil, err
	}
	if danmuAuth == nil && err == nil {
		danmuAuth.Buid = uint(req.Buid)
		danmuAuth.UUID = req.Client_uuid
		err = l.svcCtx.DanmuAuthDB.Save(l.ctx, danmuAuth)
		if err != nil {
			return nil, err
		}

		new_vcode := vcode.GenBiliVCodeWithExtraInfo(danmuAuth.UUID, string(rune(danmuAuth.Buid)), danmuAuth.CreatedAt.Format("2006-01-02 15:04:05"), "vc-", 6)
		danmuAuth.VCode = new_vcode
		err = l.svcCtx.DanmuAuthDB.Save(l.ctx, danmuAuth)
		if err != nil {
			return nil, err
		}

		return &types.ApplyResponse{
			Vcode: danmuAuth.VCode,
		}, nil
	}

	return &types.ApplyResponse{
		Vcode: danmuAuth.VCode,
	}, nil
}
