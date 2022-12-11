package service

import (
	"go.uber.org/zap"
	"goreddit/dao/mysql"
	"goreddit/dao/redis"
	"goreddit/model"
	"goreddit/pkg/snowflake"
)

func CreatePost(post *model.Post) (err error) {
	post.ID = snowflake.GenID()
	err = mysql.CreatePost(post)
	if err != nil {
		return
	}
	// 把创建时间写入Redis
	err = redis.CreatePost(post.ID)
	return
}

func GetPostByID(postId int64) (data *model.ApiPostDetail, err error) {
	post, err := mysql.GetPostByID(postId)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(postId) failed", zap.Int64("pid", postId), zap.Error(err))
		return
	}

	// 根据作者ID查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}

	// 根据社区ID查询社区信息
	community, err := mysql.GetCommunityDetailById(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
		return
	}

	return &model.ApiPostDetail{
		AuthorName:      user.UserName,
		Post:            post,
		CommunityDetail: community,
	}, nil
}

func GetPostList(pageNum, pageSize int64) (data []*model.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(pageNum, pageSize)
	if err != nil {
		return
	}
	data = make([]*model.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		// 根据作者ID查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}

		// 根据社区ID查询社区信息
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue
		}
		data = append(data, &model.ApiPostDetail{
			AuthorName:      user.UserName,
			Post:            post,
			CommunityDetail: community,
		})
	}
	return
}

func GetPostList2(p *model.ParamPostList) (data []*model.ApiPostDetail, err error) {
	// 1.去Redis查Id列表
	ids, err := redis.GetPostIdsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return
	}
	posts, err := mysql.GetPostListByIds(ids)
	if err != nil {
		return nil, err
	}

	//
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	data = make([]*model.ApiPostDetail, 0, len(posts))
	for idx, post := range posts {
		// 根据作者ID查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}

		// 根据社区ID查询社区信息
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue
		}
		data = append(data, &model.ApiPostDetail{
			AuthorName:      user.UserName,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		})
	}
	return
}
