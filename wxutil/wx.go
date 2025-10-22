package wxutil

import (
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

// 获取当前时间的微信时间戳格式字符串
func GenWxTimestampNow() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// 微信支付 金额转字符串
func WxMoneyToFen(m decimal.Decimal) string {
	return m.Mul(decimal.NewFromInt(100)).String()
}

// 微信支付 金额字符串转金额
func WxFenToMoney(num string) (decimal.Decimal, error) {
	fromString, err := decimal.NewFromString(num)
	if err != nil {
		return decimal.Zero, err
	}
	return fromString.Div(decimal.NewFromInt(100)), nil
}
