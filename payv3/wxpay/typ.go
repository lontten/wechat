package wxpay

type WxPayType int

const (
	APP WxPayType = iota + 1
	MiniProgram
	Pc_Web
	Phone_Web
	Wx_Web
)
