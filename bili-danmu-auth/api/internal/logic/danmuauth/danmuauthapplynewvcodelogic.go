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

func (l *DanmuAuthApplyNewVCodeLogic) DanmuAuthApplyNewVCode(req *types.ApplyNewVCodeRequest) (*types.ApplyNewVCodeResponse, error) {
	l.Logger.Infof("DanmuAuthApplyNewVCode, req: %+v", req)

	var new_vcode string

	switch req.Key {
	case "": // devloper login or signup
		// set periodlimit for devloper login or signup
		if !l.svcCtx.Limiter.TryAcquire() {
			l.Logger.Errorf("too many requests")
			return nil, errors.New("too many requests")
		}

		new_vcode = vcode.GenRandomBiliVCode(req.ClientID, req.Buid, l.svcCtx.Config.DanmuAuth.DevloperVCodePrefix, 11)
		err := l.svcCtx.DanmuAuthDB.Save(l.ctx, &core.DanmuAuth{Buid: req.Buid, ClientID: req.ClientID, VCode: new_vcode})
		if err != nil {
			l.Logger.Errorf("save danmu auth log failed, err: %v", err)
			return nil, err
		}

		// init balance if not exist
		balc, err := l.svcCtx.BalanceDB.FindByBuid(l.ctx, req.Buid)
		if err != nil {
			l.Logger.Errorf("find balance failed, err: %v", err)
			return nil, err
		} else if balc == nil && err == nil { // if not exist, create new
			l.Logger.Infof("balance not found, init balance")
			l.svcCtx.BalanceDB.Save(l.ctx, &core.Balance{Buid: req.Buid, Balance: l.svcCtx.Config.DanmuAuth.InitialBalance})
		}

		// check if access key exist
		keys, err := l.svcCtx.AccessKeyDB.ListByBuid(l.ctx, req.Buid)
		if err != nil {
			l.Logger.Errorf("list access key failed, err: %v", err)
			return nil, err
		} else if len(keys) == 0 && err == nil { // if no access key, add a new one
			l.Logger.Infof("access key not found, add a new one")
			newAccessKey, _ := uuid.NewV4()
			l.svcCtx.AccessKeyDB.Save(l.ctx, &core.AccessKey{Buid: req.Buid, Key: newAccessKey.String()})
		}

	default: // normal user login
		// check if key is valid
		accessKey, err := l.svcCtx.AccessKeyDB.FindByKey(l.ctx, req.Key)
		if err != nil {
			l.Logger.Errorf("find access key failed, err: %v", err)
			return nil, err
		} else if accessKey == nil && err == nil {
			l.Logger.Errorf("access key not found")
			return nil, errors.New("access key not found")
		}

		balc, err := l.svcCtx.BalanceDB.FindByBuid(l.ctx, accessKey.Buid)
		if err != nil {
			l.Logger.Errorf("find balance failed, err: %v", err)
			return nil, err
		} else if balc == nil && err == nil {
			l.Logger.Errorf("balance not found, please contact admin")
			return nil, errors.New("balance not found, please contact admin")
		}

		// decr balance 1
		err = l.svcCtx.BalanceDB.DecrBalance(l.ctx, balc, 1)
		if err != nil {
			l.Logger.Errorf("decr balance failed, err: %v", err)
			return nil, err
		}

		danmu_auth_log, err := l.svcCtx.DanmuAuthDB.FindByClientIDAndBuid(l.ctx, req.ClientID, req.Buid, l.svcCtx.Config.DanmuAuth.VCodeExpire)
		if err != nil {
			l.Logger.Errorf("find danmu auth log failed, err: %v", err)
			return nil, err
		} else if danmu_auth_log == nil && err == nil { // not found, create new
			l.Logger.Infof("danmu auth log not found, create new")
			new_vcode = vcode.GenRandomBiliVCode(req.ClientID, req.Buid, l.svcCtx.Config.DanmuAuth.NormalUserVCodePrefix, 10)
			l.svcCtx.DanmuAuthDB.Save(l.ctx, &core.DanmuAuth{Buid: req.Buid, ClientID: req.ClientID, VCode: new_vcode})
		} else {
			new_vcode = danmu_auth_log.VCode
		}
	}

	return &types.ApplyNewVCodeResponse{Vcode: new_vcode, ClientID: req.ClientID}, nil

}
