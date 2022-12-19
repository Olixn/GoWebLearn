package router

import (
	"net/http"
	"time"

	"github.com/Olixn/GoWebLearn/controller"

	"github.com/Olixn/GoWebLearn/middleware"

	"github.com/Olixn/GoWebLearn/setting"

	"github.com/Olixn/GoWebLearn/logger"
	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.RateLimitMiddleware(2*time.Second, 1))

	v1 := r.Group("/api/v1")
	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	// 登录业务路由
	v1.POST("/login", controller.LoginHandler)
	// 使用中间件
	v1.Use(middleware.JWTAuthMiddleware())
	{
		// 社区分类列表业务路由
		v1.GET("/community", controller.CommunityHandler)
		// 分类详情
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		// 发布帖子
		v1.POST("/post", controller.CreatePostHandler)
		// 帖子详情
		v1.GET("/post/:id", controller.PostDetailHandler)
		v1.GET("/posts", controller.PostListHandler)
		v1.GET("/posts2", controller.PostListHandler2)

		// 投票
		v1.POST("/vote", controller.PostVoteHandler)
		v1.GET("/version", func(c *gin.Context) {
			c.String(http.StatusOK, setting.Conf.Version)
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404 not found",
		})
	})
	return r
}
