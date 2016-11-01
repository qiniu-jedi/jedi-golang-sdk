# jedi-golang-sdk

本版本适用于 go 1.7.* 
更低版本暂时未测

---

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

