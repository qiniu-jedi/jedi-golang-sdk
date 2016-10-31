package jedi

import (
	"testing"
)

func TestGetVideoInfo(t *testing.T) {

	if _, err := GetVideoInfo("if-dianbo", "fish1.mp4"); err != nil {
		t.Error("获取视频信息 \nGetVideoInfo() Faied ")
	}
}

func TestGetVideoList(t *testing.T) {
	var c ConfQiniu
	c.SetMac("ak", "sk")
	if _, err := GetVideoList(c, "if-dianbo", "", 2); err != nil {
		t.Error("获取视频信息列表 \nGetVideoList() Faied ")
	}
}
