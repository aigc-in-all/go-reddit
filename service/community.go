package service

import (
	"goreddit/dao/mysql"
	"goreddit/model"
)

func GetCommunityList() (list []*model.Community, err error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetailById(id int64) (data *model.CommunityDetail, err error) {
	return mysql.GetCommunityDetailById(id)
}
