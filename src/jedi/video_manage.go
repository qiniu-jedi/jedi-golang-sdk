package jedi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// updateInfo 视频信息
type updateInfo struct {
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
}

/*
hub 点播空间名
key 视频key
name 视频name
*/

// GetVideoInfo 获取视频信息
//<Hub>:点播空间名
//<key>:  对video key(对应源bucket key)做 urlsafe的base64编码
func GetVideoInfo(hub, key string) (resData VideoItem, err error) {
	//base64 encoded url
	encodedKey := base64.URLEncoding.EncodeToString([]byte(key))
	urlStr := fmt.Sprintf("%s/v1/hubs/%s/videos/%s", QINIU_JEDI_HOST, hub, encodedKey)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		log.Println(err)
		return resData, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed get data from api:\n", err.Error())
		return resData, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read from responese\n", err)
		return resData, err
	}
	json.Unmarshal(data, &resData)
	return resData, nil
}

// GetVideoList 获取视频信息列表
func GetVideoList(c ConfQiniu, hub, cursor string, count int) (resData videoItems, err error) {
	var urlStr string
	//第一次不用提供cursor
	if cursor == "" {
		urlStr = fmt.Sprintf("%s/v1/hubs/%s/videos?count=%d", QINIU_JEDI_HOST, hub, count)
	} else {
		urlStr = fmt.Sprintf("%s/v1/hubs/%s/videos?cursor=%s&count=%d", QINIU_JEDI_HOST, hub, cursor, count)
	}

	data, err := RequestWithoutBody("GET", urlStr, c)
	if err != nil {
		return resData, err
	}

	json.Unmarshal(data, &resData)
	return resData, nil
}

// UpdateVideoInfo 更新视频信息
func UpdateVideoInfo(c ConfQiniu, hub, key, name, description string, tags []string) (res string, err error) {

	updateinfo := updateInfo{name, tags, description}
	urlStr := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/videos/" + base64.URLEncoding.EncodeToString([]byte(key))
	resData, err := RequestWithBody("PUT", urlStr, updateinfo, c)
	if err != nil {
		return err.Error(), err
	}
	return string(resData), nil
}

// SetVideoImage 设置封面
// active 封面索引
func SetVideoImage(c ConfQiniu, hub, key string, active int) (res string, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/videos/%s/thumbnails/active/%d",
		QINIU_JEDI_HOST, hub, base64.URLEncoding.EncodeToString([]byte(key)), active)
	resData, err := RequestWithoutBody("PUT", urlStr, c)
	if err != nil {
		return err.Error(), err
	}
	return string(resData), nil

}

// DeleteVideo 删除视频
func DeleteVideo(c ConfQiniu, hub, key string) (res string, err error) {

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/videos/%s",
		QINIU_JEDI_HOST, hub, base64.URLEncoding.EncodeToString([]byte(key)))
	resData, err := RequestWithoutBody("DELETE", urlStr, c)
	if err != nil {
		return err.Error(), err
	}
	return string(resData), nil
}

// BatchdeleteVideos  批量删除
// videoKeys 视频列表
func BatchdeleteVideos(c ConfQiniu, hub string, videoKeys []string) (res DeleteInfo, err error) {

	for i, v := range videoKeys {
		videoKeys[i] = base64.URLEncoding.EncodeToString([]byte(v))
	}
	//构造请求body
	keys := struct {
		Keys []string `json:"keys"`
	}{videoKeys}

	urlStr := fmt.Sprintf("%s/v1/hubs/%s/videos", QINIU_JEDI_HOST, hub)
	data, err := RequestWithBody("DELETE", urlStr, keys, c)
	if err != nil {
		return res, err
	}
	json.Unmarshal(data, &res)
	return res, nil

}
