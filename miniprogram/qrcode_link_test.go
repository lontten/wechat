package miniprogram

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMiniProgramConfig_GetQRCode(t *testing.T) {
	var data = GetUnlimitedQRCodeReq{
		Scene:      "fadsfdas",
		Path:       "aa",
		CheckPath:  nil,
		Width:      0,
		AutoColor:  false,
		LineColor:  nil,
		IsHyaline:  false,
		EnvVersion: "",
	}
	jsonBody, _ := json.Marshal(data)
	fmt.Println(string(jsonBody))
}
