package errors

const (
	ConfigErrorCode = 40001
	RespCodeErrorCode
)

var (
	ConfigError   = New(ConfigErrorCode, "请遵循配置说明！")
	RespCodeError = New(RespCodeErrorCode, "返回错误")
)
