package mysql

import (
	"github.com/jmoiron/sqlx"
	"goreddit/model"
	"strings"
)

func CreatePost(post *model.Post) (err error) {
	sqlStr := `insert into post (
                  post_id, title, content, author_id, community_id)
                  values (?, ?, ?, ?, ?)
                  `
	_, err = db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	return
}

func GetPostByID(postId int64) (post *model.Post, err error) {
	sqlStr := "select post_id, title, content, author_id, community_id, create_time from post where post_id = ?"
	post = new(model.Post)
	err = db.Get(post, sqlStr, postId)
	return
}

func GetPostList(pageNum, pageSize int64) (posts []*model.Post, err error) {
	sqlStr := `select 
    post_id, title, content, author_id, community_id, create_time 
	from post 
	order by create_time desc 
	limit ?,?`
	posts = make([]*model.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (pageNum-1)*pageSize, pageSize)
	return
}

func GetPostListByIds(ids []string) (posts []*model.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post 
	where post_id in (?) 
	order by FIND_IN_SET(post_id, ?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&posts, query, args...)
	return
}
