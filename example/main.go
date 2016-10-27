package main

import (
	// "encoding/base64"
	"fmt"
	"jedi"
)

func main() {
	jedi.SetMac("ak", "sk")

	//获取上传凭证√
	// res := jedi.GetUpToken("if-dianbo", "", 86400)

	//上传测试√
	// res := jedi.UploadVideoFile("if-dianbo", "", "/Users/tianyang/Downloads/last.mkv", "last.mkv", 86400)

	//查询视频√
	// res := jedi.GetVideoInfo("if-dianbo", "fish1.mp4")

	//批量查询 第一次 √
	// res := jedi.GetVideoList("if-dianbo", "", 2)
	//批量查询 第二次 √
	// res := jedi.GetVideoList("if-dianbo", "MTQ3NzQ5MTgxODcxMDAwMDAwMA==", 2)

	//更新视频√
	// res := jedi.UpdateVideoInfo("if-dianbo", "java-deploy.mp4", "hahahaha.mp4", "description 3", []string{"xxxx", "b", "c", "d"})

	//制定封面√
	// res := jedi.SetVideoImage("if-dianbo", "fish2.mp4", 3)

	//删除视频 √
	// res := jedi.DeleteVideo("if-dianbo", "fish2.mp4")

	//批量删除 √
	// res := jedi.BatchdeleteVideos("if-dianbo", []string{"fish3.mp4", "fish4.mp4"})

	//创建转码配置 √ >>"id":"5811c3c164703ca43300002c"
	// res := jedi.CreateConf("if-dianbo", "conf2")

	//删除转码配置 √
	// res := jedi.DeleteConf("if-dianbo", "5811b9dd64703c4c2f0003d7")

	// videopreset := jedi.VideoInfo{"150k", 1280, 720, "24", "libx264"}
	// audiopreset := jedi.AudioInfo{"256k", 50, "44100", "libfaac"}
	// preset := jedi.TransferConf{"preset2", false, "mp4", videopreset, audiopreset}
	//创建转码预设 √  5811c3df64703ca43300002e
	// res := jedi.CreatePreset("if-dianbo", "5811c3c164703ca43300002c", preset)

	// GetPreset 获取转码预设√
	// res := jedi.GetPreset("if-dianbo", "5811c3c164703ca43300002c", "5811c3df64703ca43300002e")

	// UpgradePreset 更新转码预设 √
	// res := jedi.UpgradePreset("if-dianbo", "5811c3c164703ca43300002c", "5811c3df64703ca43300002e", preset)

	//删除转码预设 √
	// res := jedi.DeletePreset("if-dianbo", "5811c3c164703ca43300002c", "5811c3cc64703ca43300002d")

	//启用禁用转码预设 √
	res := jedi.EnablePreset("if-dianbo", "5811c3c164703ca43300002c", "5811c3df64703ca43300002e", "0")

	//获取所有转码配置 √
	// res := jedi.ListTransConf("if-dianbo")
	fmt.Printf("%s\n", string(res))
}
