package jedi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// QINIU_JEDI_HOST API域名
const QINIU_JEDI_HOST = "http://jedi.qiniuapi.com"

//  ACCESS_KEY qiniu access key
var ACCESS_KEY string

// SECRET_KEY qiniu secret key
var SECRET_KEY string

// SetMac 初始化ak,sk
func SetMac(accessKey, secretKey string) {
	ACCESS_KEY, SECRET_KEY = accessKey, secretKey
}

//Sign return the QiniuToken for Authorization
//Every request to jedi API must Add a Header Authorization which value is tis token
func Sign(req *http.Request) (token string) {
	h := hmac.New(sha1.New, []byte(SECRET_KEY))
	url := req.URL
	//add path
	data := req.Method + " " + url.Path
	io.WriteString(h, data)

	//add query if it's not null
	if url.RawQuery != "" {

		io.WriteString(h, "?"+url.RawQuery)
	}
	//add host
	io.WriteString(h, "\nHost: "+req.Host)

	//Add Content-Type
	conType := req.Header.Get("Content-Type")
	if conType != "" {
		io.WriteString(h, "\nContent-Type: "+conType)
	}

	io.WriteString(h, "\n\n")

	var body []byte
	if req.Body != nil {
		body, _ = ioutil.ReadAll(req.Body)
	} else {
		body = []byte("")
	}
	if string(body) != "" && conType != "" &&
		conType != "application/octet-stream" {
		//Add Body into data
		h.Write(body)
	}

	tmpToken := h.Sum(nil)
	token = "Qiniu " + ACCESS_KEY + ":" + base64.URLEncoding.EncodeToString(tmpToken)
	return

}

//RequestWithoutBody 对api发出请求并且返回response body
func RequestWithoutBody(method, url string) []byte {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("Failed Make a Request", err)
		return []byte(err.Error())
	}
	token := Sign(req)
	req.Header.Add("Authorization", token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("", err)
		return []byte(err.Error())
	}

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read from responese\n", err)
		return []byte(err.Error())
	}
	if string(resData) == "{}" {
		return []byte("OK")
	}
	return resData
}

//RequestWithBody 带body对api发出请求并且返回response body
func RequestWithBody(method, url string, body interface{}) []byte {
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Println("json error", err)
		return []byte(err.Error())
	}

	reqForToken, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		log.Println("Req Token:", err)
		return []byte(err.Error())
	}
	reqForToken.Header.Set("Content-Type", "application/json")
	token := Sign(reqForToken)

	updateReq, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		log.Println(err)
		return []byte(err.Error())
	}
	updateReq.Header.Set("Authorization", token)
	updateReq.Header.Set("Content-Type", "application/json")
	client := http.Client{}

	resp, err := client.Do(updateReq)
	log.Println(1)

	if err != nil {
		log.Println("Failed get data from api:\n", err.Error())
		return []byte(err.Error())
	}
	defer resp.Body.Close()

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read from responese\n", err)
		return []byte(err.Error())
	}
	if string(resData) == "{}" {
		return []byte("OK")
	}
	return resData
}
