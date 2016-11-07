package jedi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type ConfQiniu struct {
	AccessKey string
	SecretKey string
}

// SetMac 初始化ak,sk
func (c *ConfQiniu) SetMac(accessKey, secretKey string) {
	c.AccessKey, c.SecretKey = accessKey, secretKey
}

//Sign return the QiniuToken for Authorization
//Every request to jedi API must Add a Header Authorization which value is this token
func (c *ConfQiniu) Sign(req *http.Request) string {

	//Get a sha1 HMAC hash with secretkey
	h := hmac.New(sha1.New, []byte(c.SecretKey))

	url := req.URL

	//add path
	data := req.Method + " " + url.Path
	io.WriteString(h, data)

	//add query if it's not null
	if url.RawQuery != "" {

		io.WriteString(h, "?"+url.RawQuery)
	}

	//add Host
	io.WriteString(h, "\nHost: "+req.Host)

	//Add Content-Type
	conType := req.Header.Get("Content-Type")
	if conType != "" {
		io.WriteString(h, "\nContent-Type: "+conType)
	}

	// Add Return
	io.WriteString(h, "\n\n")

	var body []byte
	//Add body if satisfy conditions
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

	// calculate HMAC-SHA1 signature
	tmpToken := h.Sum(nil)
	token := fmt.Sprintf("Qiniu %s:%s", c.AccessKey, base64.URLEncoding.EncodeToString(tmpToken))

	return token

}

// RequestWithoutBody 对api发出请求并且返回response body
// 该函数为无须 body的request通用函数
func RequestWithoutBody(method, url string, c ConfQiniu) (resData []byte, err error) {

	// Make a new request with corresponding method and url
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("Failed Make a Request", err)
		return []byte(err.Error()), err
	}
	//get a qiniu token and Add it into header
	token := c.Sign(req)
	req.Header.Add("Authorization", token)

	// Send out request and read response
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("", err)
		return []byte(err.Error()), err
	}

	resData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read from responese\n", err)
		return []byte(err.Error()), err
	}

	if resp.StatusCode != 200 {
		return resData, errors.New(string(resData))
	}

	if string(resData) == "{}" {
		return []byte("OK"), nil
	}
	return resData, nil
}

//RequestWithBody 带body对api发出请求并且返回response body
func RequestWithBody(method, url string, body interface{}, c ConfQiniu) (resData []byte, err error) {

	//Get a json data to post
	reqBody, err := json.Marshal(body)
	if err != nil {
		// log.Println("json error", err)
		return []byte(err.Error()), err
	}
	//// Make a new request with corresponding method, url and body
	reqForToken, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		// log.Println("Req Token:", err)
		return []byte(err.Error()), err
	}
	reqForToken.Header.Set("Content-Type", "application/json")
	//get qiniu token
	token := c.Sign(reqForToken)

	updateReq, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		// log.Println(err)
		return []byte(err.Error()), err
	}
	updateReq.Header.Set("Authorization", token)
	updateReq.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	resp, err := client.Do(updateReq)

	if err != nil {
		// log.Println("Failed get data from api:\n", err.Error())
		return []byte(err.Error()), err
	}
	defer resp.Body.Close()
	// fmt.Println(resp.StatusCode)

	resData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// log.Println("Failed to read from responese\n", err)
		return []byte(err.Error()), err
	}

	if resp.StatusCode != 200 {
		return resData, errors.New(string(resData))
	}

	if string(resData) == "{}" {
		return []byte("OK"), nil
	}
	return resData, nil
}
