package logic

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"os"

	"picture/internal/svc"
	"picture/internal/types"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiveImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReceiveImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiveImgLogic {
	return &ReceiveImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReceiveImgLogic) ReceiveImg(req *types.ImgReq) (resp *types.ImgRes, err error) {

	resp = &types.ImgRes{Code: 200, Message: "success"}
	os.Mkdir("./file", os.FileMode(0777))
	uid := "./file/" + uuid.NewString()
	content, _ := json.Marshal(req)
	_ = os.WriteFile(uid+"req.json", content, os.FileMode(0777))
	_ = os.WriteFile(uid+"req.chassisPhoto", []byte(req.ChassisPhoto), os.FileMode(0777))
	_ = os.WriteFile(uid+"req.vehiclePhoto", []byte(req.VehiclePhoto), os.FileMode(0777))
	if req.ChassisPhoto != "" {
		imgContent, err := base64.StdEncoding.DecodeString(req.ChassisPhoto)
		if err == nil {
			_ = os.WriteFile(uid+"ChassisPhoto.png", imgContent, os.FileMode(0777))
		}
	}
	if req.VehiclePhoto != "" {
		imgContent, err := base64.StdEncoding.DecodeString(req.VehiclePhoto)
		if err == nil {
			_ = os.WriteFile(uid+"VehiclePhoto.png", imgContent, os.FileMode(0777))
		}
	}
	l.Logger.Infof("uid:%s,req_size=%d", uid, len(content))
	return
}
