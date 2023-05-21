package devloper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthDelKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthDelKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthDelKeyLogic {
	return &DanmuAuthDelKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthDelKeyLogic) DanmuAuthDelKey(req *types.DeleteKeyRequest) (resp *types.DeleteKeyResponse, err error) {
	l.Logger.Infof("DanmuAuthDelKey,req: %v", req)

	var buid_string string = fmt.Sprintf("%v", l.ctx.Value("buid"))

	buid, err := strconv.Atoi(buid_string)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.AccessKeyDB.Delete(l.ctx, buid)
	if err != nil {
		return nil, err
	}

	return &types.DeleteKeyResponse{}, nil
}
