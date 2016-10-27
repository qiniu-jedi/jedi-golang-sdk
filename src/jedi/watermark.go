package jedi

import "strconv"

// WatMakConf 水印配置
type WatMakConf struct {
	Type      string `json:"type"`
	Enable    bool   `json:"enabled"`
	Gravity   string `json:"gravity"`
	OffsetX   int    `json:"offset_x"`
	OffsetY   int    `json:"offset_Y"`
	Image     string `json:"image"`
	Text      string `json:"text"`
	TextFont  string `json:"text_font"`
	TextColor string `json:"text_color"`
	TextSize  int    `json:"text_size"`
}

// CrtWatMak 创建水印配置
func CrtWatMak(hub string, conf WatMakConf) string {
	url := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/watermarks"
	res := RequestWithBody("POST", url, conf)
	return string(res)
}

// UpgradeWatConf 更新(覆盖)水印配置
func UpgradeWatConf(hub string, WatMakID int, conf WatMakConf) string {
	url := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/watermarks/" + strconv.Itoa(WatMakID)
	res := RequestWithBody("PUT", url, conf)
	return string(res)
}

// DelWatMakConf 删除水印配置
func DelWatMakConf(hub string, WatMakID int) string {
	url := QINIU_JEDI_HOST + "/v1/" + hub + "/watermarks/" + strconv.Itoa(WatMakID)
	res := RequestWithoutBody("DELETE", url)
	return string(res)
}

// GetWatConf 获取水印配置
func GetWatConf(hub string, watMakID int) string {
	url := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/watermarks/" + strconv.Itoa(watMakID)
	res := RequestWithoutBody("GET", url)
	return string(res)
}

// EnableWatMak 启用禁用水印配置
func EnableWatMak(hub string, watMakID int, enable bool) string {
	url := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/watermarks/" +
		strconv.Itoa(watMakID) + "/enabled/" + strconv.FormatBool(enable)

	res := RequestWithoutBody("PUT", url)
	return string(res)
}
