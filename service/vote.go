package service

import (
	"go.uber.org/zap"
	"goreddit/dao/redis"
	"goreddit/model"
	"strconv"
)

// VoteForPost 为帖子投票
func VoteForPost(userId int64, p *model.ParamVoteData) error {
	zap.L().Debug("service.VoteForPost",
		zap.Int64("userId", userId),
		zap.String("postId", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userId)), p.PostID, float64(p.Direction))
}
