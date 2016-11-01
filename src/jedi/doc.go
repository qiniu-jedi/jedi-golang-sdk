/*
使用方法
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


*/

package jedi
