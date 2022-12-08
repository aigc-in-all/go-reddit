package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goreddit/service"
	"strconv"
)

func CommunityListHandler(c *gin.Context) {
	list, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("service.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, list)
}

func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := service.GetCommunityDetailById(id)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
