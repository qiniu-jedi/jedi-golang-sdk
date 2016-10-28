package jedi

import (
	"context"
	"encoding/json"

	"qiniupkg.com/api.v7/kodocli"
)

// UploadTokenBody 上传参数
type UploadTokenBody struct {
	Hub      string `json:"hub"`
	DeadLine int    `json:"deadline"`
	Type     string `json:"type"`
}

// UploadVideoFile 上传文件
//Hub  点播空间名
//videoType  video类型
//filePath  文件目录
//key 保存文件名
//deadline token有效时间
func UploadVideoFile(hub, videoType, filePath, key string, deadline int) error {

	zone := 0
	uploader := kodocli.NewUploader(zone, nil)
	ctx := context.Background()

	//Get Up Token
	uploadToken := GetUpToken(hub, videoType, deadline)

	//token结构题
	var uptoken struct {
		Uptoken string `json:"uptoken"`
	}
	err := json.Unmarshal([]byte(uploadToken), &uptoken)
	if err != nil {
		return err
	}
	err = uploader.PutFile(ctx, nil, uptoken.Uptoken, key, filePath, nil)
	if err != nil {
		return err
	}
	return nil

}

//GetUpToken 获取uptoken
//Hub  点播空间名
//deadline  有效时间
//aksk  用户密钥
func GetUpToken(hub, videoType string, deadline int) string {

	if videoType == "" {
		videoType = "video"
	}

	urlStr := QINIU_JEDI_HOST + "/v1/uptokens"
	body := UploadTokenBody{hub, deadline, videoType}
	//获取uptoken
	data := RequestWithBody("POST", urlStr, body)

	return string(data)

}
