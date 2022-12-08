package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"goreddit/model"
)

func CheckUserExist(userName string) (err error) {
	sqlStr := "select count(user_id) from user where username=?"
	var count int
	if err := db.Get(&count, sqlStr, userName); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

func InsertUser(user *model.User) (err error) {
	// 对密码进行加密
	password := encryptPassword(user.Password)
	sqlStr := "insert into user(user_id, username, password) values (?, ?, ?)"
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, password)
	return
}

func Login(user *model.User) (err error) {
	encryptPassword := encryptPassword(user.Password)
	sqlStr := "select user_id, username, password from user where username=?"
	if err := db.Get(user, sqlStr, user.UserName); err != nil {
		if err == sql.ErrNoRows {
			return ErrorUserNotExist
		}
		return err
	}
	if user.Password != encryptPassword {
		return ErrorInvalidPassword
	}
	return nil
}

func encryptPassword(input string) string {
	h := md5.New()
	h.Write([]byte("goreddit2022"))
	return hex.EncodeToString(h.Sum([]byte(input)))
}
