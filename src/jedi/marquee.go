package jedi

import (
	"encoding/json"
	"fmt"
)

func CreateUpdateMarquee(c ConfQiniu, hub string, conf MarqueeConf) (res string, err error) {
	urlStr := fmt.Sprintf("%s/v1/hubs/%s/marquee", QINIU_JEDI_HOST, hub)
	resData, err := RequestWithBody("POST", urlStr, conf, c)
	if err != nil {
		return res, err
	}
	//json.Unmarshal(resData, &res)
	return string(resData), nil
}

func GetMarqueeConf(c ConfQiniu, hub string) (res MarqueeConf, err error) {
	urlStr := fmt.Sprintf("%s/v1/hubs/%s", QINIU_JEDI_HOST, hub)
	resData, err := RequestWithoutBody("GET", urlStr, c)
	if err != nil {
		return res, err
	}
	json.Unmarshal(resData, &res)
	return res, nil
}

func EnableMarquee(c ConfQiniu, hub string, enable bool) (res string, err error) {
	var urlStr string
	if enable {
		urlStr = fmt.Sprintf("%s/v1/hubs/%s/enabled/1", QINIU_JEDI_HOST, hub)
	} else {
		urlStr = fmt.Sprintf("%s/v1/hubs/%s/enabled/0", QINIU_JEDI_HOST, hub)
	}
	resData, err := RequestWithoutBody("PUT", urlStr, c)
	if err != nil {
		return res, err
	}
	return string(resData), nil
}
