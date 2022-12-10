package mysql

import "goreddit/model"

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
	limit ?,?`
	posts = make([]*model.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (pageNum-1)*pageSize, pageSize)
	return
}
