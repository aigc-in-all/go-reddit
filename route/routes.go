package route

import (
	"github.com/gin-gonic/gin"
	"goreddit/logger"
	"goreddit/setting"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, setting.Conf.Version)
	})
	return r
}
