package logic

import (
	"context"
	"fmt"

	"picture/internal/svc"
	"picture/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PositionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PositionLogic {
	return &PositionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PositionLogic) Position(req *types.PositionReq) (resp *types.PositionRes, err error) {
	l.Logger.Infof("position:%+v", req)
	resp = &types.PositionRes{
		Code:    200,
		Message: "ok" + fmt.Sprintf("%+v", req),
	}
	return
}
