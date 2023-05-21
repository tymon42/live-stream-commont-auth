package devloper

import (
	"context"
	"fmt"
	"strconv"

	// "github.com/fox-one/pkg/uuid"
	"github.com/gofrs/uuid"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthAddKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthAddKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthAddKeyLogic {
	return &DanmuAuthAddKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthAddKeyLogic) DanmuAuthAddKey(req *types.AddKeyRequest) (*types.AddKeyResponse, error) {
	l.Logger.Infof("DanmuAuthAddKey,req: %v", req)

	var buid_string string = fmt.Sprintf("%v", l.ctx.Value("buid"))

	buid, err := strconv.Atoi(buid_string)
	if err != nil {
		return nil, err
	}

	newKey, _ := uuid.NewV4()

	err = l.svcCtx.AccessKeyDB.Save(l.ctx, &core.AccessKey{Buid: buid, Key: newKey.String()})
	if err != nil {
		return nil, err
	}

	return &types.AddKeyResponse{Key: newKey.String()}, nil
}
