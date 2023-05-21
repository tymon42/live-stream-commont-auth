package devloper

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthGetKeyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthGetKeyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthGetKeyListLogic {
	return &DanmuAuthGetKeyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthGetKeyListLogic) DanmuAuthGetKeyList(req *types.GetKeyListRequest) (*types.GetKeyListResponse, error) {
	l.Logger.Infof("DanmuAuthGetKeyList,req: %v", req)

	var buid_string string = fmt.Sprintf("%v", l.ctx.Value("buid"))

	buid, err := strconv.Atoi(buid_string)
	if err != nil {
		return nil, err
	}

	AccessKeys, err := l.svcCtx.AccessKeyDB.ListByBuid(l.ctx, buid)
	if err != nil {
		return nil, err
	} else if len(AccessKeys) == 0 && err == nil {
		return nil, errors.New("keys not found")
	}

	var keys []string

	// convert keys to types.GetKeyListResponse
	for _, key := range AccessKeys {
		keys = append(keys, key.Key)
	}

	// query balance
	var balc int
	balance, err := l.svcCtx.BalanceDB.FindByBuid(l.ctx, buid)
	if err != nil {
		return nil, err
	} else if balance == nil && err == nil {
		balc = 0
	} else {
		balc = balance.Balance
	}

	return &types.GetKeyListResponse{Keys: keys, Balance: balc}, nil
}
