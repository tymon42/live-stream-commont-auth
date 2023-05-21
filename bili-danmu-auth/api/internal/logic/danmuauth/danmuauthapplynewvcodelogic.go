package danmuauth

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
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

	// devloper login or signup
	if req.Key == "" {
		new_vcode := vcode.GenRandomBiliVCode(req.ClientID, req.Buid, "开发者登录或注册-", 11)
		err := l.svcCtx.DanmuAuthDB.Save(l.ctx, &core.DanmuAuth{Buid: req.Buid, ClientID: req.ClientID, VCode: new_vcode})
		if err != nil {
			return nil, err
		}

		// init balance if not exist
		balc, err := l.svcCtx.BalanceDB.FindByBuid(l.ctx, req.Buid)
		if err != nil {
			return nil, err
		} else if balc == nil && err == nil {
			err = l.svcCtx.BalanceDB.Save(l.ctx, &core.Balance{Buid: req.Buid, Balance: 50})
			if err != nil {
				return nil, err
			}
		}

		// if no access key, add a new one
		keys, err := l.svcCtx.AccessKeyDB.ListByBuid(l.ctx, req.Buid)
		if err != nil {
			return nil, err
		} else if len(keys) == 0 && err == nil {
			// add a new access key
			newAccessKey, _ := uuid.NewV4()
			err = l.svcCtx.AccessKeyDB.Save(l.ctx, &core.AccessKey{Buid: req.Buid, Key: newAccessKey.String()})
			if err != nil {
				return nil, err
			}
		}

		resp = &types.ApplyNewVCodeResponse{Vcode: new_vcode}

		return resp, nil
	}

	// check if key is valid
	accessKey, err := l.svcCtx.AccessKeyDB.FindByKey(l.ctx, req.Key)
	if err != nil {
		return nil, err
	} else if accessKey == nil && err == nil {
		return nil, errors.New("access key not found")
	}

	balc, err := l.svcCtx.BalanceDB.FindByBuid(l.ctx, accessKey.Buid)
	if err != nil {
		return nil, err
	} else if balc == nil && err == nil {
		return nil, errors.New("balance not found")
	}
	// decr balance 1
	err = l.svcCtx.BalanceDB.DecrBalance(l.ctx, balc, 1)
	if err != nil {
		return nil, err
	}

	da, err := l.svcCtx.DanmuAuthDB.FindByClientID(l.ctx, req.ClientID)
	if err != nil {
		return nil, err
	} else if da == nil && err == nil { // not found, create new
		// generate new vcode
		new_vcode := vcode.GenRandomBiliVCode(req.ClientID, req.Buid, "", 10)
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
