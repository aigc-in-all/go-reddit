package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"goreddit/model"
)

func CheckUserExist(userName string) (bool, error) {
	sqlStr := "select count(user_id) from user where username=?"
	var count int
	if err := db.Get(&count, sqlStr, userName); err != nil {
		return false, err
	}
	return count > 0, nil
}

func InsertUser(user *model.User) (err error) {
	// 对密码进行加密
	password := encryptPassword(user.Password)
	sqlStr := "insert into user(user_id, username, password) values (?, ?, ?)"
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, password)
	return
}

func encryptPassword(input string) string {
	h := md5.New()
	h.Write([]byte("goreddit2022"))
	return hex.EncodeToString(h.Sum([]byte(input)))
}
