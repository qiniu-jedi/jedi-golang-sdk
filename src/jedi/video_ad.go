package jedi

import (
	"encoding/json"
	"fmt"
)

func CreateAd(c ConfQiniu, hub string, conf VideoAdConf) (res AdID, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/ad/videos", QINIU_JEDI_HOST, hub)
	resData, err := RequestWithBody("POST", urlStr, conf, c)
	if err != nil {
		return res, err
	}
	json.Unmarshal(resData, &res)
	return res, nil

}

func UpadateAd(c ConfQiniu, hub, VAdID string, conf VideoAdConf) (res string, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/ad/videos/%s", QINIU_JEDI_HOST, hub, VAdID)
	resData, err := RequestWithBody("PUT", urlStr, conf, c)
	if err != nil {
		return string(resData), err
	}
	return string(resData), nil
}

func QueryVideoAd(c ConfQiniu, hub, VAdID string) (res VideoAdConf, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/ad/videos/%s", QINIU_JEDI_HOST, hub, VAdID)
	resData, err := RequestWithoutBody("GET", urlStr, c)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resData, &res)
	return res, nil
}

func GetVideoAds(c ConfQiniu, hub string) (res VideoAds, err error) {
	urlStr := fmt.Sprintf("%s/v1/hubs/%s/ad/videos", QINIU_JEDI_HOST, hub)
	resData, err := RequestWithoutBody("GET", urlStr, c)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resData, &res)
	return res, nil
}

func DeleteVAd(c ConfQiniu, hub, VAdID string) (res string, err error) {
	urlStr := fmt.Sprintf("%s/v1/hubs/%s/ad/videos/%s", QINIU_JEDI_HOST, hub, VAdID)
	resData, err := RequestWithoutBody("DELETE", urlStr, c)
	if err != nil {
		return string(resData), err
	}
	return string(resData), nil
}

func BatchDeleteVAD(c ConfQiniu, hub string, ids []string) (res AdsBatchDeleteConf, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/ad/videos", QINIU_JEDI_HOST, hub)

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

func EnableVAds(c ConfQiniu, hub, VAdID string, enabled bool) (res string, err error) {

	var urlStr string
	if enabled {
		urlStr = fmt.Sprintf("%s/v1/hubs/%s/ad/videos/%s/enabled/%d", QINIU_JEDI_HOST, hub, VAdID, 1)
	} else {
		urlStr = fmt.Sprintf("%s/v1/hubs/%s/ad/videos/%s/enabled/%d", QINIU_JEDI_HOST, hub, VAdID, 0)
	}

	resData, err := RequestWithoutBody("PUT", urlStr, c)
	if err != nil {
		return string(resData), err
	}
	return string(resData), nil
}
