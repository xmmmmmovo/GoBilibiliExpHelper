package errors

const (
	// HasReceviedRewardManga 已领取漫画大会员福利
	HasReceviedRewardManga = 6
	// HasReceviedRewardHome 已领取漫画大会员福利
	HasReceviedRewardHome = 69801
	// ConfigErrorCode 配置失败状态码
	ConfigErrorCode = 40001
	// RespCodeErrorCode 返回消息失败状态码
	RespCodeErrorCode
)

var (
	// ConfigError 配置失败错误
	ConfigError = New(ConfigErrorCode, "请遵循配置说明！")
	// RespCodeError 返回消息失败错误
	RespCodeError = New(RespCodeErrorCode, "返回错误")
)
