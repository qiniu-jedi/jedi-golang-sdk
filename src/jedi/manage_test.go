package jedi

import "testing"

func TestUpdateVideoInfo(t *testing.T) {
	var c ConfQiniu
	c.SetMac("ak", "sk")
	if _, err := UpdateVideoInfo(c, "if-dianbo", "java-deploy.mp4", "hahahaha.mp4", "description 3", []string{"xxxx", "b", "c", "d"}); err != nil {
		t.Error("更新视频信息 \n UpdateVideoInfo() Faied ")
	}
}

func TestSetVideoImage(t *testing.T) {
	var c ConfQiniu
	c.SetMac("ak", "sk")
	if _, err := SetVideoImage(c, "if-dianbo", "fish2.mp4", 3); err != nil {
		t.Error("指定封面 \n SetVideoImage() Faied ")
	}
}
func TestDeleteVideo(t *testing.T) {
	var c ConfQiniu
	c.SetMac("ak", "sk")
	if _, err := DeleteVideo(c, "if-dianbo", "fish2.mp4"); err != nil {
		t.Error("删除视频 \n DeleteVideo() Faied ")
	}
}

func TestBatchdeleteVideos(t *testing.T) {
	var c ConfQiniu
	c.SetMac("ak", "sk")
	if _, err := BatchdeleteVideos(c, "if-dianbo", []string{"fish3.mp4", "fish4.mp4"}); err != nil {
		t.Error("批量删除 \n BatchdeleteVideos() Faied ")
	}
}
