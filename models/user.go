package models

type UserInfo struct {
	IsLogin       bool
	Uname         string
	VipStatus     int
	VipType       int
	CurrentLevel  int
	Money         float64
	BcoinBalance  float64
	CouponBalance float64
}

type UserLogin struct {
}
