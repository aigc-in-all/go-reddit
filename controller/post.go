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
	pageNum, pageSize := getPageInfo(c)

	data, err := service.GetPostList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("service.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按时间或分数排序查询帖子列表
// @Tag 帖子相关接口
// @Accept application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query model.ParamPostList false "参数查询"
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数 /api/v1/posts/page=1&size=10&order=time
	// 获取请求参数
	p := &model.ParamPostList{
		Page:  1,
		Size:  10,
		Order: model.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := service.GetPostList2(p)
	if err != nil {
		zap.L().Error("service.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// getPageInfo 获取分页数据
func getPageInfo(c *gin.Context) (page, size int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var err error
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return
}
