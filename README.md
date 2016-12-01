# jedi-golang-sdk

[![GoDoc](http://godoc.org/github.com/astaxie/beego?status.svg)](https://godoc.org/github.com/TianZong48/jedi-golang-sdk)

本版本适用于 go 1.7.* 
更低版本暂时未测

##Before Start
上传基于七牛存储go sdk
```bash
go get -u qiniupkg.com/api.v7
```


##使用方法：
```go
	var c jedi.ConfQiniu
	c.SetMac("your_ak", "you_sk")

	//获取uptoken
	res,err := jedi.GetUpToken(c, "hubname", "", 86400)
	if err !=nil{
		//Code
	}


	//删除视频 
	res2,err := jedi.DeleteVideo(c,"hubname", "key")
	if err !=nil{
		//Code
	}
	...

```

## API列表
#### HTTP请求鉴权

#### 资源上传

#### 视频管理* [x] 查询视频* [x] 批查询视频* [x] 更新视频* [x] 指定封面* [x] 删除视频* [x] 批量删除视频
#### 转码配置和预设* [x] 创建转码配置* [x] 删除转码配置* [x] 创建转码预设* [x] 获取转码预设* [x] 更新转码预设* [x] 删除转码预设* [x] 启用禁用转码预设* [x] 获取所有转码配置
#### 水印配置* [x]  创建水印配置* [x] 更新印配置* [x] 删除水印配置* [x] 获取水印配置* [x] 获取所有水印配置* [x] 启水禁用水印配置
 ####跑马灯配置
* [ ] 创建/ 新跑马灯配置 
* [ ] 获取跑马灯配置 
* [ ] 启用禁用跑马灯配置
####视频广告管 * [ ] 创建视频广告* [ ] 新视频广告* [ ] 获取视频广告* [ ] 获取全部视频广告* [ ] 删除视频广告* [ ] 批量删除视频广告* [ ] 启用禁用视频广告

#### 贴片广告管理 * [ ] 创建贴片广告* [ ] 更新贴片广告* [ ] 获取贴片馆告* [ ] 获取全部贴片广告* [ ] 删除贴片广告* [ ] 批量删除贴片广告* [ ] 启用禁用贴片广告
#### DRM 配置* [ ] 设置 DRM 保护模式 
* [ ] 获取视频 户加密密钥
####错误格式


