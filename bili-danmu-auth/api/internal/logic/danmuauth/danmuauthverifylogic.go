package danmuauth

import (
	"context"
	"errors"
	"time"

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

func (l *DanmuAuthVerifyLogic) DanmuAuthVerify(req *types.VerifyRequest) (*types.VerifyResponse, error) {
	l.Logger.Infof("verify request: %+v", req)

	da, err := l.svcCtx.DanmuAuthDB.FindByClientID(l.ctx, req.ClientID)
	if err != nil {
		return nil, err
	} else if da == nil && err == nil {
		return nil, errors.New("client_id not found")
	}

	if da.VerifiedCount < 1 {
		return nil, errors.New("vcode not verified")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, req.Buid, req.ClientID)
	if err != nil {
		return nil, err
	}

	resp := &types.VerifyResponse{
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
	}

	return resp, nil
}

func (l *DanmuAuthVerifyLogic) getJwtToken(secretKey string, iat, seconds int64, buid int, client_id string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["buid"] = buid
	claims["client_id"] = client_id
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
