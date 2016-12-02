package jedi

import (
	"encoding/json"
	"fmt"
)

func SetDRMSecurity(c ConfQiniu, hub string, mode int) (res string, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/protection/%d", QINIU_JEDI_HOST, hub, mode)

	resData, err := RequestWithoutBody("PUT", urlStr, c)
	if err != nil {
		return "", err
	}
	return string(resData), nil
}

func GetUKey(c ConfQiniu, hub, encodedkey string) (res drmInfo, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/videos/%s/ueky", QINIU_JEDI_HOST, hub, encodedkey)

	resData, err := RequestWithoutBody("GET", urlStr, c)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resData, &res)
	return res, nil
}
