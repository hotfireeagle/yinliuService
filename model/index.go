package model

type ResponseCode int

const (
	Ok           ResponseCode = iota + 1 // 响应正常
	Err                                  // 响应错误
	UnLogin                              // 未登录
	LoginOverdue                         // 登录过期
)

type IResponse struct {
	Code ResponseCode `json:"code"` // 响应状态
	Msg  string       `json:"msg"`  // 报错信息
	Data interface{}  `json:"data"` // 正式数据
}
