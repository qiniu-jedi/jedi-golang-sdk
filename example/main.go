package main

import (
	// "encoding/base64"
	"fmt"
	"jedi"
)

func main() {
	var c jedi.ConfQiniu
	c.SetMac("Ome3aovf0XEAl6ATGQKKpMj4jqrws1_Hsziopq1U", "818Wub4AGBUoAjwaaa51LTeycjzl-fodGn7AltqZ")

	//获取上传凭证√
	// res := jedi.GetUpToken(c, "if-dianbo", "", 86400)

	//上传测试√
	// res := jedi.UploadVideoFile(c, "if-dianbo", "", "/Users/tianyang/Desktop/自.mp4", 86400)

	//查询视频√
	// res, _ := jedi.GetVideoInfo("if-dianbo", "fish1.mp4")

	//批量查询 第一次 √
	// res, _ := jedi.GetVideoList(c, "if-dianbo", "", 2)
	//批量查询 第二次 √
	// res := jedi.GetVideoList(c, "if-dianbo", "MTQ3Nzg5ODU1OTcyMDAwMDAwMA==", 2)

	//更新视频√
	// res, _ := jedi.UpdateVideoInfo(c, "if-dianbo", "java-deploy.mp4", "hahahaha.mp4", "description 3", []string{"xxxx", "b", "c", "d"})

	//制定封面√
	// res := jedi.SetVideoImage(c,"if-dianbo", "fish2.mp4", 3)

	//删除视频 √
	// res := jedi.DeleteVideo(c,"if-dianbo", "fish2.mp4")

	//批量删除 √
	// res := jedi.BatchdeleteVideos(c,"if-dianbo", []string{"fish3.mp4", "fish4.mp4"})

	//创建转码配置 √
	// res, _ := jedi.CreateConf(c, "if-dianbo", "conf2")

	//删除转码配置 √
	// res, _ := jedi.DeleteConf(c, "if-dianbo", "581717cdcaf65587db0002dd")

	// videopreset := jedi.VideoInfo{"150k", 1280, 720, "24", "libx264"}
	// audiopreset := jedi.AudioInfo{"256k", 50, "44100", "libfaac"}
	// preset := jedi.TransferConf{"preset2", false, "mp4", videopreset, audiopreset}
	//创建转码预设 √
	// res := jedi.CreatePreset(c,"if-dianbo", "", preset)

	// GetPreset 获取转码预设√
	// res := jedi.GetPreset(c,"if-dianbo", "", "")

	// UpgradePreset 更新转码预设 √
	// res := jedi.UpgradePreset(c,"if-dianbo", "", "", preset)

	//删除转码预设 √
	// res := jedi.DeletePreset(c,"if-dianbo", "", "")

	//启用禁用转码预设 √
	// res := jedi.EnablePreset(c,"if-dianbo", "", "", "0")

	//获取所有转码配置 √
	// res, _ := jedi.ListTransConf(c, "if-dianbo")
	fmt.Printf("%s\n", string(res))
}
