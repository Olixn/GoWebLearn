package router

import (
	"net/http"

	"github.com/Olixn/GoWebLearn/controller"

	"github.com/Olixn/GoWebLearn/setting"

	"github.com/Olixn/GoWebLearn/logger"
	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, setting.Conf.Version)
	})

	return r
}
