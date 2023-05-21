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

type DanmuAuthCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthCheckLogic {
	return &DanmuAuthCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthCheckLogic) DanmuAuthCheck(req *types.CheckRequest) (*types.CheckResponse, error) {
	var buid_string string = fmt.Sprintf("%v", l.ctx.Value("buid"))

	buid, err := strconv.Atoi(buid_string)
	if err != nil {
		return nil, err
	}

	// buid is not nil, not 0, not empty
	if l.ctx.Value("buid") == nil || buid == 0 {
		return nil, errors.New("buid is empty")
	}

	return &types.CheckResponse{Buid: buid}, nil
}
