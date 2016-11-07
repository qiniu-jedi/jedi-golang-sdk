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

