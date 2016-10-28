package jedi

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// GetVideoInfo 获取视频信息
//<Hub>:点播空间名
//<key>:  对video key(对应源bucket key)做 urlsafe的base64编码
func GetVideoInfo(hub, key string) string {
	//base64 encoded url
	encodedKey := base64.URLEncoding.EncodeToString([]byte(key))
	urlStr := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/videos/" + encodedKey
	// fmt.Println(urlStr)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed get data from api:\n", err.Error())
		return err.Error()
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read from responese\n", err)
		return err.Error()
	}

	return string(data)
}

// GetVideoList 获取视频信息列表
func GetVideoList(hub, cursor string, count int) string {
	var urlStr string
	//第一次不用提供cursor
	if cursor == "" {
		urlStr = QINIU_JEDI_HOST + "/v1/hubs/" + hub +
			"/videos?count=" + strconv.Itoa(count)
	} else {
		urlStr = QINIU_JEDI_HOST + "/v1/hubs/" + hub +
			"/videos?cursor=" +
			cursor +
			"&count=" + strconv.Itoa(count)
	}

	data := RequestWithoutBody("GET", urlStr)

	return string(data)
}
