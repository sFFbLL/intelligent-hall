package api

import (
	"github.com/gin-gonic/gin"
	"go_api/settings"
	"go_api/utils/tcpServe"

	"go_api/app/dto"
	"go_api/utils/app"
)

func Power(c *gin.Context) {
	// 1.获取参数 校验参数
	var power dto.LightDto
	if err := c.ShouldBind(&power); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2. TCP 发送命令
	var tcpErr error
	if power.Name == "cabinet" && *power.Switch == 1 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.PowerOpen)
	} else if power.Name == "cabinet" && *power.Switch == 0 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.PowerClose)
	} else if power.Name == "tv" && *power.Switch == 1 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.TvOpen)
	} else if power.Name == "tv" && *power.Switch == 0 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.TvClose)
	} else if power.Name == "lamplight" && *power.Switch == 1 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.LightAllOpen)
	} else if power.Name == "lamplight" && *power.Switch == 0 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.LightAllClose)
	} else if power.Name == "dvd" && *power.Switch == 1 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdOpen)
	} else if power.Name == "dvd" && *power.Switch == 0 {
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdClose)
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
