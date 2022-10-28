package api

import (
	"github.com/gin-gonic/gin"
	"go_api/app/dto"
	"go_api/settings"
	"go_api/utils/app"
	"go_api/utils/tcpServe"
)

func Dvd(c *gin.Context) {
	// 1.获取参数 校验参数
	var dvdDto dto.DvdDto
	if err := c.ShouldBind(&dvdDto); err != nil {
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}

	// 2. TCP 发送命令
	var tcpErr error
	switch dvdDto.Status {
	case "Power":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdPower)
	case "Setting":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdSetting)
	case "ChannelAdd":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdChannelAdd)
	case "ChannelSub":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdChannelSub)
	case "Up":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdUp)
	case "Down":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdDown)
	case "Left":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdLeft)
	case "Right":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdRight)
	case "Ok":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdOk)
	case "VolumeAdd":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdVolumeAdd)
	case "VolumeSub":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdVolumeSub)
	case "Return":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdReturn)
	case "Retreat":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdRetreat)
	case "HomeMenu":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdHomeMenu)
	case "Advance":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdAdvance)
	case "MediaCenter":
		tcpErr = tcpServe.TcpDate(settings.CommandData.DvdMediaCenter)
	default:
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
