package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"goreddit/model"
)

func GetCommunityList() (list []*model.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&list, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailById(id int64) (data *model.CommunityDetail, err error) {
	sqlStr := "select community_id, community_name, introduction, create_time from community where community_id = ?"
	data = new(model.CommunityDetail)
	if err := db.Get(data, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("there is no community detail in db")
			err = ErrorInvalidID
		}
	}
	return data, err
}
