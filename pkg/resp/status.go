package resp

type ErrorStatus struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

type ErrorStatus400 struct {
	Code    uint32 `json:"code" example:"3"`
	Message string `json:"message" example:"参数错误"`
}

type ErrorStatus401 struct {
	Code    uint32 `json:"code" example:"16"`
	Message string `json:"message" example:"未授权"`
}

type ErrorStatus403 struct {
	Code    uint32 `json:"code" example:"7"`
	Message string `json:"message" example:"禁止操作"`
}

type ErrorStatus404 struct {
	Code    uint32 `json:"code" example:"5"`
	Message string `json:"message" example:"资源未找到"`
}

type ErrorStatus500 struct {
	Code    uint32 `json:"code" example:"10"`
	Message string `json:"message" example:"服务器错误"`
}
