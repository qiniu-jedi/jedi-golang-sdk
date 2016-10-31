package jedi

import "testing"

func TestCreateConf(t *testing.T) {
	var c ConfQiniu
	c.SetMac("ak", "sk")
	if _, err := CreateConf(c, "if-dianbo", "conf2"); err != nil {
		t.Error("创建转码配置 \n CreateConf() Faied ")
	}
}

func TestDeleteConf(t *testing.T) {
	var c ConfQiniu
	c.SetMac("ak", "sk")
	if _, err := DeleteConf(c, "if-dianbo", "581717cdcaf65587db0002dd"); err != nil {
		t.Error("删除转码配置 \n DeleteConf() Faied ")
	}
}
func TestCreatePreset(t *testing.T) {
	var c ConfQiniu
	c.SetMac("ak", "sk")

	videopreset := VideoInfo{"150k", 1280, 720, "24", "libx264"}
	audiopreset := AudioInfo{"256k", 50, "44100", "libfaac"}
	preset := TransferConf{"preset2", false, "mp4", videopreset, audiopreset}

	if _, err := CreatePreset(c, "if-dianbo", "", preset); err != nil {
		t.Error("创建转码预设 \n CreatePreset() Faied ")
	}
}

func TestGetPreset(t *testing.T) {
	var c ConfQiniu
	c.SetMac("ak", "sk")
	videopreset := VideoInfo{"150k", 1280, 720, "24", "libx264"}
	audiopreset := AudioInfo{"256k", 50, "44100", "libfaac"}
	preset := TransferConf{"preset2", false, "mp4", videopreset, audiopreset}

	if _, err := UpgradePreset(c, "if-dianbo", "", "", preset); err != nil {
		t.Error("更新转码预设 \n GetPreset() Faied ")
	}
}

func TestEnablePreset(t *testing.T) {
	var c ConfQiniu
	c.SetMac("ak", "sk")
	if _, err := EnablePreset(c, "if-dianbo", "", "", "0"); err != nil {
		t.Error("启用禁用转码预设 \n EnablePreset() Faied ")
	}
}
func TestListTransConf(t *testing.T) {
	var c ConfQiniu
	c.SetMac("ak", "sk")
	if _, err := ListTransConf(c, "if-dianbo"); err != nil {
		t.Error("获取所有转码配置 \n ListTransConf() Faied ")
	}
}
