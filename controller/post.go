package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goreddit/model"
	"goreddit/service"
	"strconv"
)

func CreatePostHandler(c *gin.Context) {
	p := new(model.Post)
	if err := c.ShouldBind(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	if err := service.CreatePost(p); err != nil {
		zap.L().Error("service.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	post, err := service.GetPostByID(pid)
	if err != nil {
		zap.L().Error("service.GetPostByID failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}

func GetPostListHandler(c *gin.Context) {
	// 获取分页数据
	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageSize")
	pageNum, err := strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		pageNum = 0
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageNum = 10
	}

	data, err := service.GetPostList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("service.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
