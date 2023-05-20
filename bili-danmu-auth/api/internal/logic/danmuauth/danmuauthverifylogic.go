package danmuauth

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/svc"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DanmuAuthVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDanmuAuthVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DanmuAuthVerifyLogic {
	return &DanmuAuthVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DanmuAuthVerifyLogic) DanmuAuthVerify(req *types.VerifyRequest) (resp *types.VerifyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

func (l *DanmuAuthVerifyLogic) getJwtToken(secretKey string, iat, seconds, buid uint64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["buid"] = buid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
