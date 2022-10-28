package utils

import (
	"bytes"
	"encoding/json"
	"go_api/app/config"
	"go_api/app/dto"
	"go_api/settings"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 登录获取token
func GetHttpToken() (err error) {
	result, err := HttpPost(config.ApplicationConfig.TerminalHttpLogin, settings.LoginData, "application/json")
	var login dto.Login
	if err = json.Unmarshal(result, &login); err != nil {
		return
	}
	settings.Token = login.Token
	settings.TerminalList.Token = login.Token
	return
}

func Get(url string) (dataString string, err error) {
	// 发送GET请求
	// url：         请求地址
	// response：    请求返回的内容
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, errRead := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if errRead != nil && errRead == io.EOF {
			break
		} else if errRead != nil {
			err = errRead
			return
		}
	}
	dataString = result.String()
	return
}

func HttpPost(url string, data interface{}, contentType string) (dataString []byte, err error) {
	// 发送POST请求
	// url：         请求地址
	// data：        POST请求提交的数据
	// contentType： 请求体格式，如：application/json
	// content：     请求放回的内容
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dataString, _ = ioutil.ReadAll(resp.Body)
	return
}
