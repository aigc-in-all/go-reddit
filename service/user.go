package service

import (
	"goreddit/dao/mysql"
	"goreddit/model"
	"goreddit/pkg/jwt"
	"goreddit/pkg/snowflake"
)

func SignUp(p model.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	err = mysql.CheckUserExist(p.UserName)
	if err != nil {
		return err
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

func Login(p *model.ParamLogin) (token string, err error) {
	user := &model.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	token, err = jwt.GenToken(user.UserID, user.UserName)
	return
}
