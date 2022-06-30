package proto

type ReportRequest struct {
	Action  string  `json:"action"`
	Payload string  `json:"payload"`
	Runtime Runtime `json:"runtime"`
}

type Runtime struct {
	IP              string `json:"ip"`
	DeviceID        string `json:"device_id"`
	Platform        string `json:"platform"`         //平台类型(比如群晖)
	PlatformVersion string `json:"platform_version"` //平台版本(比如6.1;7.1)
	ClientVersion   string `json:"client_version"`   //客户端(b下载器)的版本
	OS              string `json:"os"`
}

type ReportResponse struct {
	Ret int64  `json:"ret"`
	Msg string `json:"msg"`
}
