package jedi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
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
