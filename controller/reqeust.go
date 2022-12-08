package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var CtxUserIDKey = "userId"

var ErrorUserNotLogin = errors.New("user not login")

// getCurrentUser 获取当前登录用户的ID
func getCurrentUser(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
