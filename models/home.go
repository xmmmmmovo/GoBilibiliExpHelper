package models

// DailyRewardStatus 每日奖励情况
type DailyRewardStatus struct {
	Login bool `json:"login"` // 今日登录奖励
	Watch bool `json:"watch"` // 观看奖励
	Coins int  `json:"coins"` // 投币奖励
	Share bool `json:"share"` // 分享奖励
}
