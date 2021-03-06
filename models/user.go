package models

// UserInfo 用户信息结构体
type UserInfo struct {
	IsLogin       bool    // 是否登录
	Uname         string  // 用户名
	VipStatus     int     // 是否有大会员
	VipType       int     // 大会员类型 1"月度 2:年度
	CurrentLevel  int     // 当前等级
	Money         float64 // 硬币数
	BcoinBalance  float64 // b币数
	CouponBalance float64 // 每月奖励b币数
}

// UserLogin 用户登录结构体
type UserLogin struct {
}

// VipPrivilege 大会员权益列表
type VipPrivilege struct {
	Type       int `json:"type"`
	State      int `json:"state"`
	ExpireTime int `json:"expire_time"`
}
