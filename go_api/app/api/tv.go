package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_api/app/config"
	"go_api/app/dto"
	"go_api/settings"
	"go_api/utils"
	"go_api/utils/app"
	"strconv"
	"strings"
)

// tv中控
func TvIsPlay(c *gin.Context) {
	// 1.获取参数 校验参数
	var isPlay dto.IsPlay
	if err := c.ShouldBind(&isPlay); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2. 发送指令
	var status string
	if *isPlay.Status == 1 {
		status = "start"
	} else {
		status = "stop"
	}
	for _, data := range isPlay.Data {
		dataString, err := utils.Get("http://" + data.Ip + config.ApplicationConfig.TerminalOperateList + "?" +
			"id=" + strconv.Itoa(data.ShowId) + "&operation=" + status + "&offset=0")
		if err != nil {
			app.ResponseError(c, app.CodeParamNotComplete)
			return
		}
		var data map[string]interface{}
		if err = json.Unmarshal([]byte(dataString), &data); err != nil {
			app.ResponseError(c, app.CodeParamNotComplete)
			return
		}
		if data["code"] != float64(200) {
			app.ResponseError(c, app.CodeParamNotComplete)
			return
		}
	}

	// 3.返回数据
	app.ResponseSuccess(c, nil)
}

// tv列表
func TvList(c *gin.Context) {
	// 1. 获取数据
	dataString, err := utils.HttpPost(config.ApplicationConfig.TerminalHttpList,
		settings.TerminalList, "application/json")
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	datalist := string(dataString)
	datalist = strings.Replace(datalist, "\"{", "{", -1)
	datalist = strings.Replace(datalist, "}\"", "}", -1)

	var data map[string]interface{}
	if err = json.Unmarshal([]byte(datalist), &data); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	if data["rescode"] != float64(200) {
		_ = utils.GetHttpToken()
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	// 2.返回数据
	app.ResponseSuccess(c, data)
}

func TvQuery(c *gin.Context) {
	// 1. 校验参数
	ip := c.Query("ip")
	if len(ip) == 0 {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2. 发送请求
	dataString, err := utils.Get("http://" + ip + config.ApplicationConfig.TerminalQueryList)
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal([]byte(dataString), &data); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 3.返回数据
	if data["code"] != float64(200) {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	app.ResponseSuccess(c, data["data"])
}

func TvOperate(c *gin.Context) {
	// 1.获取参数 校验参数
	var operate dto.Operate
	if err := c.ShouldBind(&operate); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	if len(settings.OperateStatus[operate.Operation]) == 0 {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2. 发送请求
	dataString, err := utils.Get("http://" + operate.Ip + config.ApplicationConfig.TerminalOperateList + "?" +
		"id=" + strconv.Itoa(operate.ShowId) + "&operation=" + operate.Operation + "&offset=" + strconv.Itoa(operate.Offset))
	if err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal([]byte(dataString), &data); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 3.返回数据
	if data["code"] != float64(200) {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	app.ResponseSuccess(c, nil)
}
