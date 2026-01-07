package miniprogram

import "github.com/lontten/lcore/v2/netutil"

type GetQRCodeReq struct {
	Path       string `json:"path"`                  // 是	扫码进入的小程序页面路径，最大长度 1024 个字符，不能为空
	Width      int    `json:"width,omitempty"`       // 否	二维码的宽度，单位 px。默认值为430，最小 280px，最大 1280px
	AutoColor  bool   `json:"auto_color"`            // 否	默认值false；自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	LineColor  *Color `json:"line_color,omitempty"`  // 否	默认值{"r":0,"g":0,"b":0} ；auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
	IsHyaline  bool   `json:"is_hyaline"`            // 否	默认值false；是否需要透明底色，为 true 时，生成透明底色的小程序码
	EnvVersion string `json:"env_version,omitempty"` // 否	要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"。默认是正式版
}

type Color struct {
	R int `json:"r"` // 0-255
	G int `json:"g"` // 0-255
	B int `json:"b"` // 0-255
}

type GetQRCodeResp struct {
	ErrCode int    `json:"errcode"` // 错误码，请求失败时返回, 0 表示成功
	ErrMsg  string `json:"errmsg"`  // 错误信息，请求失败时返回

	Buffer []byte `json:"buffer"` // 二维码图片的二进制数据
}

// GetQRCode 获取小程序码, createQRCode 的上位替代
// 永久有效，有数量限制
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/qr-code/api_getqrcode.html
func (c MiniProgramConfig) GetQRCode(req GetQRCodeReq) (GetQRCodeResp, error) {
	url := "https://api.weixin.qq.com/wxa/getwxacode"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return GetQRCodeResp{}, err
	}
	url += "?access_token=" + accessToken

	return netutil.PostJson[GetQRCodeResp](url, req)
}

type CreateQRCodeReq struct {
	Path  string `json:"path"`            // 是	扫码进入的小程序页面路径，最大长度 128 个字符，不能为空
	Width int    `json:"width,omitempty"` // 否	二维码的宽度，单位 px。默认值为430，最小 280px，最大 1280px
}

// CreateQRCode 获取小程序二维码
// 不推荐使用，除非是为了兼容老系统；现在建议用 GetQRCode
// 永久有效，有数量限制
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/qr-code/api_createqrcode.html
func (c MiniProgramConfig) CreateQRCode(req CreateQRCodeReq) (GetQRCodeResp, error) {
	url := "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return GetQRCodeResp{}, err
	}
	url += "?access_token=" + accessToken

	return netutil.PostJson[GetQRCodeResp](url, req)
}

type GetUnlimitedQRCodeReq struct {
	Scene     string `json:"scene"`                // 是	最大32个可见字符，只支持数字，大小写英文以及部分特殊字符
	Path      string `json:"path,omitempty"`       // 否	默认是主页，页面 page，例如 pages/index/index
	CheckPath *bool  `json:"check_path,omitempty"` // 否	默认是true，检查page 是否存在，为 true 时 page 必须是已经发布的小程序存在的页面（否则报错）；为 false 时允许小程序未发布或者 page 不存在， 但page 有数量上限（60000个）请勿滥用。

	Width      int    `json:"width,omitempty"`       // 否	二维码的宽度，单位 px。默认值为430，最小 280px，最大 1280px
	AutoColor  bool   `json:"auto_color"`            // 否	默认值false；自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	LineColor  *Color `json:"line_color,omitempty"`  // 否	默认值{"r":0,"g":0,"b":0} ；auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
	IsHyaline  bool   `json:"is_hyaline"`            // 否	默认值false；是否需要透明底色，为 true 时，生成透明底色的小程序码
	EnvVersion string `json:"env_version,omitempty"` // 否	要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"。默认是正式版
}

// GetUnlimitedQRCode 获取不限制的小程序码
// 永久有效，数量暂无限制
// https://developers.weixin.qq.com/miniprogram/dev/server/API/qrcode-link/qr-code/api_getunlimitedqrcode.html
func (c MiniProgramConfig) GetUnlimitedQRCode(req GetUnlimitedQRCodeReq) (GetQRCodeResp, error) {
	url := "https://api.weixin.qq.com/wxa/getwxacode"
	accessToken, err := c.GetAccessTokenCache()
	if err != nil {
		return GetQRCodeResp{}, err
	}
	url += "?access_token=" + accessToken

	return netutil.PostJson[GetQRCodeResp](url, req)
}
