package jedi

// TransferConf 预设配置
type TransferConf struct {
	Name    string    `json:"name"`
	Enabled bool      `json:"enabled"`
	Format  string    `json:"format"`
	Video   VideoInfo `json:"video"`
	Audio   AudioInfo `json:"audio"`
}

// VideoInfo 视频参数
type VideoInfo struct {
	BitRate   string `json:"bit_rate"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	FrameRate string `json:"frame_rate"`
	CodecName string `json:"codec_name"`
}

// AudioInfo 音频参数
type AudioInfo struct {
	BitRate    string `json:"bit_rate"`
	Quality    int    `json:"quality"`
	SampleRate string `json:"sample_rate"`
	CodecName  string `json:"codec_name"`
}

// CreateConf 创建转吗配置
func CreateConf(hub, name string) string {
	url := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/transconfs"

	reqBody := struct {
		Name string `json:"name"`
	}{name}
	resData := RequestWithBody("POST", url, reqBody)
	return string(resData)
}

// DeleteConf 删除转码配置
func DeleteConf(hub, transID string) string {
	url := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/transconfs/" + transID
	resData := RequestWithoutBody("DELETE", url)
	return string(resData)
}

// CreatePreset 创建转码预设
func CreatePreset(hub, transID string, config TransferConf) string {

	urlStr := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/transconfs/" + transID + "/transsets"

	resData := RequestWithBody("POST", urlStr, config)
	return string(resData)
}

// GetPreset 获取转码预设
func GetPreset(hub, transID, transSetID string) string {
	url := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/transconfs/" + transID + "/transsets/" + transSetID
	res := RequestWithoutBody("GET", url)

	return string(res)

}

// UpgradePreset 更新转码预设
func UpgradePreset(hub, transID, transSetID string, config TransferConf) string {

	urlStr := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/transconfs/" + transID + "/transsets/" + transSetID
	resData := RequestWithBody("PUT", urlStr, config)
	return string(resData)
}

// DeletePreset 删除转码预设
func DeletePreset(hub, transID, transSetID string) string {
	url := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/transconfs/" + transID + "/transsets/" + transSetID
	resData := RequestWithoutBody("DELETE", url)
	return string(resData)
}

// EnablePreset 启用禁用转码预设
func EnablePreset(hub, transID, transSetID, enable string) []byte {
	url := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/transconfs/" + transID + "/transsets/" + transSetID + "/enabled/" + enable
	resData := RequestWithoutBody("PUT", url)
	return resData
}

// ListTransConf 获取所有转码配置
func ListTransConf(hub string) string {
	url := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/transconfsall"
	resData := RequestWithoutBody("GET", url)
	return string(resData)
}
