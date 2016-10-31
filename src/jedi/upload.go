package jedi

import (
	"context"
	"encoding/json"
	"fmt"

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
func UploadVideoFile(c ConfQiniu, hub, videoType, filePath string, deadline int) (res string, err error) {

	zone := 0
	uploader := kodocli.NewUploader(zone, nil)
	ctx := context.Background()

	//Get Up Token
	uploadToken, err := GetUpToken(c, hub, videoType, deadline)
	if err != nil {
		return err.Error(), err
	}

	//token结构题
	var uptoken struct {
		Uptoken string `json:"uptoken"`
	}
	err = json.Unmarshal([]byte(uploadToken), &uptoken)
	if err != nil {
		return err.Error(), err
	}
	err = uploader.PutFileWithoutKey(ctx, nil, uptoken.Uptoken, filePath, nil)
	if err != nil {
		return err.Error(), err
	}
	return "upload Succeed", nil

}

//GetUpToken 获取uptoken
//Hub  点播空间名
//deadline  有效时间
//aksk  用户密钥
func GetUpToken(c ConfQiniu, hub, videoType string, deadline int) (res string, err error) {

	//set type equals "video" while It's emtpy
	if videoType == "" {
		videoType = "video"
	}

	urlStr := fmt.Sprintf("%s/v1/uptokens", QINIU_JEDI_HOST)
	body := UploadTokenBody{hub, deadline, videoType}
	//获取uptoken
	resData, err := RequestWithBody("POST", urlStr, body, c)
	if err != nil {
		return err.Error(), err
	}

	return string(resData), nil

}
