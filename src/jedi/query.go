package jedi

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetVideoInfo 获取视频信息
//<Hub>:点播空间名
//<key>:  对video key(对应源bucket key)做 urlsafe的base64编码
func GetVideoInfo(hub, key string) (resData string, err error) {
	//base64 encoded url
	encodedKey := base64.URLEncoding.EncodeToString([]byte(key))
	urlStr := fmt.Sprintf("%s/v1/hubs/%s/videos/%s", QINIU_JEDI_HOST, hub, encodedKey)
	fmt.Println(urlStr)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		log.Println(err)
		return err.Error(), err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed get data from api:\n", err.Error())
		return err.Error(), err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read from responese\n", err)
		return err.Error(), err
	}

	return string(data), nil
}

// GetVideoList 获取视频信息列表
func GetVideoList(c ConfQiniu, hub, cursor string, count int) (resData string, err error) {
	var urlStr string
	//第一次不用提供cursor
	if cursor == "" {
		urlStr = fmt.Sprintf("%s/v1/hubs/%s/videos?count=%d", QINIU_JEDI_HOST, hub, count)
	} else {
		urlStr = fmt.Sprintf("%s/v1/hubs/%s/videos?cursor=%s&count=%d", QINIU_JEDI_HOST, hub, cursor, count)
	}

	data, err := RequestWithoutBody("GET", urlStr, c)

	return string(data), nil
}
