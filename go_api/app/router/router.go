package router

import (
	"go_api/app/api"
	"go_api/app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setup 路由设置
func Setup(cfg *config.Application) *gin.Engine {
	if cfg.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	v1 := r.Group("/api")
	{
		// 电源、灯光
		v1.POST("/power", api.Power)
		v1.POST("/lamplight", api.Lamplight)
		// 电视屏幕
		v1.POST("/tv/isplay", api.TvIsPlay)
		v1.GET("/tv/list", api.TvList)
		v1.GET("/tv/query", api.TvQuery)
		v1.POST("/tv/operate", api.TvOperate)
		// DVD
		v1.POST("/dvd", api.Dvd)
	}
	return r
}
