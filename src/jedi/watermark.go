package jedi

import (
	"encoding/json"
	"fmt"
)

// CreateWatrMark 创建水印配置
func CreateWatrMark(c ConfQiniu, hub string, conf WaterMarkConf) (res WaterMarkCrt, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/watermarks", QINIU_JEDI_HOST, hub)
	resData, err := RequestWithBody("POST", urlStr, conf, c)
	if err != nil {
		return res, err
	}
	json.Unmarshal(resData, &res)
	return res, nil
}

//UpdateWaterMark 更新水印配置
//   watermarkID 水印ID
func UpdateWaterMark(c ConfQiniu, hub, watermarkID string, conf WaterMarkConf) (res string, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/watermarks/%s", QINIU_JEDI_HOST, hub, watermarkID)
	resData, err := RequestWithBody("PUT", urlStr, conf, c)
	if err != nil {
		return err.Error(), err
	}

	return string(resData), nil
}

//DeleteWatermarkConf 删除水印配置
func DeleteWatermarkConf(c ConfQiniu, hub, watermarkID string) (res string, err error) {
	urlStr := fmt.Sprintf("%s/v1/hubs/%s/watermarks/%s", QINIU_JEDI_HOST, hub, watermarkID)
	resData, err := RequestWithoutBody("DELETE", urlStr, c)
	if err != nil {
		return err.Error(), err
	}

	return string(resData), nil
}

// GetWatermarkConf 获取水印配置
func GetWatermarkConf(c ConfQiniu, hub, watermarkID string) (res WaterMarkConf, err error) {
	urlStr := fmt.Sprintf("%s/v1/hubs/%s/watermarks/%s", QINIU_JEDI_HOST, hub, watermarkID)
	resData, err := RequestWithoutBody("GET", urlStr, c)
	if err != nil {
		return res, err
	}
	json.Unmarshal(resData, &res)
	return res, nil
}

// GetAllWaterConfs 获取所有水印配置
func GetAllWaterConfs(c ConfQiniu, hub string) (res WaterMarkGroup, err error) {
	urlStr := fmt.Sprintf("%s/v1/hubs/%s/watermarks", QINIU_JEDI_HOST, hub)
	resData, err := RequestWithoutBody("GET", urlStr, c)
	if err != nil {
		return res, err
	}
	json.Unmarshal(resData, &res)

	return res, nil
}

// EnableWaterMark 启动禁用水印配置
// enabeld true启用 false禁用
func EnableWaterMark(c ConfQiniu, hub, watermarkID string, enabled bool) (res string, err error) {
	var urlStr string
	if enabled {
		urlStr = fmt.Sprintf("%s/v1/hubs/%s/watermarks/%s/enabled/%d", QINIU_JEDI_HOST, hub, watermarkID, 1)
	} else {
		urlStr = fmt.Sprintf("%s/v1/hubs/%s/watermarks/%s/enabled/%d", QINIU_JEDI_HOST, hub, watermarkID, 0)
	}
	resData, err := RequestWithoutBody("PUT", urlStr, c)
	if err != nil {
		return res, err
	}

	return string(resData), nil
}
