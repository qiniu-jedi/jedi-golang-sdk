/*
本SDK是七牛点播云解决方案的Golang SDK，利用本SDK，客户可以方便地将七牛提供的点播云服务集成到自己的业务系统汇总。
视频基本信息
点播云主要解决的是视频的上传，转码，加水印等功能。其中关于视频，我们需要了解相关的一些概念。在七牛点播云系统中，一个视频的基本属性有 key ，name，tags，description，size，creation_time，modification_time，分为代表以下含义：

使用方法：
```go
	var c jedi.ConfQiniu
	c.SetMac("your_ak", "you_sk")

	//获取uptoken
	res := jedi.GetUpToken(c, "hubname", "", 86400)

```

属性	描述
key	视频文件的key，该key在指定的点播空间中唯一存在
name	视频文件的名称，该name在指定的点播空间中可以不唯一
tags	视频文件的标签，可以为一个视频指定多个标签
description	视频文件的描述
size	视频文件的大小，单位字节
creation_time	视频文件的创建时间
modification_time	视频文件的最后修改时间

视频格式信息
视频的格式信息主要是指视频文件中音频数据和视频数据的码率，编码格式，分辨率等信息，主要有如下一些内容：

format
属性	描述
	duration	视频时长，单位秒
	format_long_name	视频格式名称，一般为用户可读
	audio

属性	描述
	bit_rate	音频码率
	codec_name	音频编码格式
	sample_rate	音频采样率
	video

属性	描述
	bit_rate	视频码率
	codec_name	视频编码格式
	avg_frame_rate	视频平均帧率
	width	视频分辨率宽
	height	视频分辨率高
*/
package jedi
