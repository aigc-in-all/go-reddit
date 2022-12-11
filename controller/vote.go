package controller

import (
	"github.com/gin-gonic/gin"
	"goreddit/model"
	"goreddit/service"
)

func PostVoteHandler(c *gin.Context) {
	p := new(model.ParamVoteData)
	if err := c.ShouldBind(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	userId, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	if err := service.VoteForPost(userId, p); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
