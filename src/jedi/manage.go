package jedi

import (
	"encoding/base64"
	"strconv"
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
func UpdateVideoInfo(hub, key, name, description string, tags []string) string {

	updateinfo := updateInfo{name, tags, description}
	urlStr := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/videos/" + base64.URLEncoding.EncodeToString([]byte(key))

	res := RequestWithBody("PUT", urlStr, updateinfo)

	return string(res)
}

// SetVideoImage 设置封面
// active 封面索引
func SetVideoImage(hub, key string, active int) string {

	urlStr := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/videos/" +
		base64.URLEncoding.EncodeToString([]byte(key)) +
		"/thumbnails/active/" + strconv.Itoa(active)

	res := RequestWithoutBody("PUT", urlStr)

	return string(res)

}

// DeleteVideo 删除视频
func DeleteVideo(hub, key string) string {

	urlStr := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/videos/" +
		base64.URLEncoding.EncodeToString([]byte(key))

	res := RequestWithoutBody("DELETE", urlStr)

	return string(res)
}

// BatchdeleteVideos  批量删除
// videoKeys 视频列表
func BatchdeleteVideos(hub string, videoKeys []string) string {

	urlStr := QINIU_JEDI_HOST + "/v1/hubs/" + hub + "/videos"

	for i, v := range videoKeys {
		videoKeys[i] = base64.URLEncoding.EncodeToString([]byte(v))
	}
	//构造请求body
	keys := struct {
		Keys []string `json:"keys"`
	}{videoKeys}

	resData := RequestWithBody("DELETE", urlStr, keys)
	return string(resData)

}
