package logic

import (
	"context"

	"backend/bsc_server/api/internal/svc"
	"backend/bsc_server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiLogic {
	return &ApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiLogic) Api(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
