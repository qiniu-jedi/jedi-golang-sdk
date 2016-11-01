package jedi

import (
	"fmt"
)

// CreateConf 创建转码配置
func CreateConf(c ConfQiniu, hub, name string) (res string, err error) {
	url := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/transconfs"

	reqBody := struct {
		Name string `json:"name"`
	}{name}
	resData, err := RequestWithBody("POST", url, reqBody, c)
	if err != nil {
		return err.Error(), err
	}

	return string(resData), nil
}

// DeleteConf 删除转码配置
func DeleteConf(c ConfQiniu, hub, transID string) (res string, err error) {

	url := fmt.Sprintf("%s/v1/hubs/%s/transconfs/%s", QINIU_JEDI_HOST, hub, transID)
	resData, err := RequestWithoutBody("DELETE", url, c)
	if err != nil {
		return err.Error(), err
	}

	return string(resData), nil
}

// CreatePreset 创建转码预设
// transID 转码配置ID
func CreatePreset(c ConfQiniu, hub, transID string, config TransferConf) (res string, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/transconfs/%s/transsets", QINIU_JEDI_HOST, hub, transID)
	resData, err := RequestWithBody("POST", urlStr, config, c)
	if err != nil {
		return err.Error(), err
	}

	return string(resData), nil
}

// GetPreset 获取转码预设
// transID 转码配置ID
// transSetId 转码预设ID
func GetPreset(c ConfQiniu, hub, transID, transSetID string) (res string, err error) {

	url := fmt.Sprintf("%s/v1/hubs/%s/transconfs/%s/transsets/%s", QINIU_JEDI_HOST, hub, transID, transSetID)
	resData, err := RequestWithoutBody("GET", url, c)
	if err != nil {
		return err.Error(), err
	}

	return string(resData), nil
}

// UpgradePreset 更新转码预设
// transID 转码配置ID
// transSetId 转码预设ID
func UpgradePreset(c ConfQiniu, hub, transID, transSetID string, config TransferConf) (res string, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/transconfs/%s/transsets/%s", QINIU_JEDI_HOST, hub, transID, transSetID)
	resData, err := RequestWithBody("PUT", urlStr, config, c)
	if err != nil {
		return err.Error(), err
	}

	return string(resData), nil
}

// DeletePreset 删除转码预设
// transID 转码配置ID
// transSetId 转码预设ID
func DeletePreset(c ConfQiniu, hub, transID, transSetID string) (res string, err error) {

	url := fmt.Sprintf("%s/v1/hubs/%s/transconfs/%s/transsets/%s", QINIU_JEDI_HOST, hub, transID, transSetID)
	resData, err := RequestWithoutBody("DELETE", url, c)
	if err != nil {
		return err.Error(), err
	}

	return string(resData), nil
}

// EnablePreset 启用禁用转码预设
// transID 转码配置ID
// transSetId 转码预设ID
func EnablePreset(c ConfQiniu, hub, transID, transSetID, enable string) (res string, err error) {

	url := fmt.Sprintf("%s/v1/hubs/%s/transconfs/%s/transsets/%s/enabled/%s", QINIU_JEDI_HOST, hub, transID, transSetID, enable)
	resData, err := RequestWithoutBody("PUT", url, c)
	if err != nil {
		return err.Error(), err
	}

	return string(resData), nil
}

// ListTransConf 获取所有转码配置
func ListTransConf(c ConfQiniu, hub string) (res string, err error) {

	url := fmt.Sprintf("%s/v1/hubs/%s/transconfsall", QINIU_JEDI_HOST, hub)
	resData, err := RequestWithoutBody("GET", url, c)
	if err != nil {
		return err.Error(), err
	}

	return string(resData), nil
}
