package api

import (
	"github.com/gin-gonic/gin"
	"go_api/settings"
	"go_api/utils/tcpServe"

	"go_api/app/dto"
	"go_api/utils/app"
)

func Lamplight(c *gin.Context) {
	// 1.获取参数 校验参数
	var light dto.LightDto
	if err := c.ShouldBind(&light); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2. TCP 发送命令
	var tcpErr error
	if light.Name == "light1" && *light.Switch == 1 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.Light1Open)
	} else if light.Name == "light1" && *light.Switch == 0 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.Light1Close)
	} else if light.Name == "spotlight1" && *light.Switch == 1 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.Spotlight1Open)
	} else if light.Name == "spotlight1" && *light.Switch == 0 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.Spotlight1Close)
	} else if light.Name == "lightstrip1" && *light.Switch == 1 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.LightStrip1Open)
	} else if light.Name == "lightstrip1" && *light.Switch == 0 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.LightStrip1Close)
	} else if light.Name == "light2" && *light.Switch == 1 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.Light2Open)
	} else if light.Name == "light2" && *light.Switch == 0 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.Light2Close)
	} else if light.Name == "spotlight2" && *light.Switch == 1 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.Spotlight2Open)
	} else if light.Name == "spotlight2" && *light.Switch == 0 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.Spotlight2Close)
	} else if light.Name == "lightstrip2" && *light.Switch == 1 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.LightStrip2Open)
	} else if light.Name == "lightstrip2" && *light.Switch == 0 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.LightStrip2Close)
	} else {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	if tcpErr != nil {
		app.ResponseError(c, app.TerminalOffLine)
		return
	}

	// 3.返回数据
	app.ResponseSuccess(c, nil)
}
