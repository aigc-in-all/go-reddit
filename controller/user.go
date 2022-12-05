package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goreddit/model"
	"goreddit/service"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	var p model.ParamSignUp
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	service.SignUp(p)
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
