package jedi

import (
	"encoding/json"
	"fmt"
)

func CreateIAd(c ConfQiniu, hub string, conf ImageAdConf) (res AdID, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/ad/images", QINIU_JEDI_HOST, hub)

	resData, err := RequestWithBody("POST", urlStr, conf, c)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resData, &res)
	return res, nil
}

func UpdateIAd(c ConfQiniu, hub string, conf ImageAdConf) (res AdID, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/ad/images", QINIU_JEDI_HOST, hub)

	resData, err := RequestWithBody("PUT", urlStr, conf, c)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resData, &res)
	return res, nil
}

func QueryIAd(c ConfQiniu, hub, IAdID string) (res ImageAdConf, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/ad/images/%s", QINIU_JEDI_HOST, hub, IAdID)

	resData, err := RequestWithoutBody("GET", urlStr, c)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resData, &res)
	return res, nil
}

func GetImageAds(c ConfQiniu, hub string) (res ImageAds, err error) {
	urlStr := fmt.Sprintf("%s/v1/hubs/%s/ad/images", QINIU_JEDI_HOST, hub)

	resData, err := RequestWithoutBody("GET", urlStr, c)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resData, &res)
	return res, nil
}


func DeleteIAd(c ConfQiniu, hub, VAdID string) (res string, err error) {
	urlStr := fmt.Sprintf("%s/v1/hubs/%s/ad/images/%s", QINIU_JEDI_HOST, hub, VAdID)
	resData, err := RequestWithoutBody("DELETE", urlStr, c)
	if err != nil {
		return string(resData), err
	}
	return string(resData), nil
}

func BatchDeleteIAD(c ConfQiniu, hub string, ids []string) (res AdsBatchDeleteConf, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/ad/images", QINIU_JEDI_HOST, hub)

	reqBody := struct {
		IDS []string `json:"ids"`
	}{ids}

	resData, err := RequestWithBody("DELETE", urlStr, reqBody, c)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resData, &res)
	return res, nil

}

func EnableIAds(c ConfQiniu, hub, IAdID string, enabled bool) (res string, err error) {

	var urlStr string
	if enabled {
		urlStr = fmt.Sprintf("%s/v1/hubs/%s/ad/images/%s/enabled/%d", QINIU_JEDI_HOST, hub, IAdID, 1)
	} else {
		urlStr = fmt.Sprintf("%s/v1/hubs/%s/ad/images/%s/enabled/%d", QINIU_JEDI_HOST, hub, IAdID, 0)
	}

	resData, err := RequestWithoutBody("PUT", urlStr, c)
	if err != nil {
		return string(resData), err
	}
	return string(resData), nil
}

