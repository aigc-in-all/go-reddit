package service

import (
	"errors"
	"goreddit/dao/mysql"
	"goreddit/model"
	"goreddit/pkg/snowflake"
)

func SignUp(p model.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	var exist bool
	exist, err = mysql.CheckUserExist(p.UserName)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("用户已存在")
	}
	// 2.构造一个User实例
	user := &model.User{
		UserID:   snowflake.GenID(),
		UserName: p.UserName,
		Password: p.Password,
	}
	// 3.保存到数据库
	return mysql.InsertUser(user)
}
