package main

import (
	// "encoding/base64"
	"fmt"
	"jedi"
)

func main() {
	var c jedi.ConfQiniu
	c.SetMac("ak", "sk")

	//获取上传凭证√
	// res := jedi.GetUpToken(c, "if-dianbo", "", 86400)

	//上传测试√
	// res, err := jedi.UploadVideoFile(c, "if-dianbo", "", "/Users/tianyang/Desktop/自定义.mp4", 86400)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//查询视频√
	// res, err := jedi.GetVideoInfo("if-dianbo", "自定义.mp4")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//批量查询 第一次 √
	// res, err := jedi.GetVideoList(c, "if-dianbo", "", 2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//批量查询 第二次 √
	// res,err := jedi.GetVideoList(c, "if-dianbo", "MTQ3Nzg5ODU1OTcyMDAwMDAwMA==", 2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	ret

	//更新视频√
	// res, err := jedi.UpdateVideoInfo(c, "if-dianbo", "java-deploy.mp4", "hahahaha.mp4", "description 3", []string{"xxxx", "b", "c", "d"})
	// if err != nil {
	// 	fmt.Println(err)
	// 	ret

	//制定封面√
	// res,err := jedi.SetVideoImage(c,"if-dianbo", "fish2.mp4", 3)
	// if err != nil {
	// 	fmt.Println(err)
	// 	ret

	//删除视频 √
	// res,err := jedi.DeleteVideo(c,"if-dianbo", "fish2.mp4")
	// if err != nil {
	// 	fmt.Println(err)
	// 	ret

	//批量删除 √
	// res, err := jedi.BatchdeleteVideos(c, "if-dianbo", []string{"fish1.mp4"})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//创建转码配置 √
	// res, err := jedi.CreateConf(c, "if-dianbo", "conf2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//删除转码配置 √
	// res, err := jedi.DeleteConf(c, "if-dianbo", "581717cdcaf65587db0002dd")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// videopreset := jedi.VideoInfo{"150k", 1280, 720, "24", "libx264"}
	// audiopreset := jedi.AudioInfo{"256k", 50, "44100", "libfaac"}
	// preset := jedi.TransferConf{"preset2", false, "mp4", videopreset, audiopreset}
	//创建转码预设 √
	// res, err := jedi.CreatePreset(c, "if-dianbo", "5818271064703ca43300038f", preset)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// GetPreset 获取转码预设√
	// res, err := jedi.GetPreset(c, "if-dianbo", "", "")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// UpgradePreset 更新转码预设 √
	// res, err := jedi.UpgradePreset(c, "", "", "", preset)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//删除转码预设 √
	// res, err := jedi.DeletePreset(c, "", "", "")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//启用禁用转码预设 √
	// res, err := jedi.EnablePreset(c, "", "", "", "0")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//获取所有转码配置 √
	// res, err := jedi.ListTransConf(c, "if-dianbo")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%v\n", res)
}
