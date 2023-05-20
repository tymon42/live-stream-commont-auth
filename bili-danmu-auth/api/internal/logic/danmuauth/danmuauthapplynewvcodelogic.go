package danmuauth

import (
	"context"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/core"
	"github.com/tymon42/live-stream-commont-auth/vcode"

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
	l.Logger.Infof("DanmuAuthApplyNewVCode, req: %+v", req)

	// TODO: check if key is valid or developer is valid

	da, err := l.svcCtx.DanmuAuthDB.FindByClientID(l.ctx, req.ClientID)
	if err != nil {
		return nil, err
	} else if da == nil && err == nil { // not found, create new
		// generate new vcode
		new_vcode := vcode.GenRandomBiliVCode(req.ClientID, string(rune(req.Buid)), "", 10)
		err := l.svcCtx.DanmuAuthDB.Save(l.ctx, &core.DanmuAuth{Buid: req.Buid, ClientID: req.ClientID, VCode: new_vcode})
		if err != nil {
			return nil, err
		}
		resp = &types.ApplyNewVCodeResponse{Vcode: new_vcode}

		return resp, nil
	}

	// found, return old vcode
	resp = &types.ApplyNewVCodeResponse{Vcode: da.VCode}

	return resp, nil
}
