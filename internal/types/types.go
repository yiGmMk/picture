// Code generated by goctl. DO NOT EDIT.
package types

type ImgReq struct {
	VehicleNum   string `json:"vehicleNum,optional"`
	GenerateTime string `json:"generateTime,optional"`
	Channel      string `json:"channel,optional"`
	ChassisPhoto string `json:"chassisPhoto,optional"`
	VehiclePhoto string `json:"vehiclePhoto,optional"`
}

type ImgRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PositionReq struct {
	DeviceType   string  `form:"device_type,optional"`
	Imei         string  `form:"imei,optional"`
	TimeBeginStr string  `form:"time_begin_str,optional"`
	TimeBegin    string  `form:"time_begin,optional"`
	IsReply      bool    `form:"is_reply,optional"`
	IsTrack      bool    `form:"is_track,optional"`
	City         string  `form:"city,optional"`
	Address      string  `form:"address,optional"`
	Lon          float64 `form:"lon,optional"`
	Lat          float64 `form:"lat,optional"`
	Type         int     `form:"type,optional"`
}

type PositionRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
