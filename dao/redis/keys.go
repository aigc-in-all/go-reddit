package redis

const (
	Prefix                 = "goreddit:"
	KeyPostTimeZSet        = "post:time"   // zset: 帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  // zset: 帖子及投票分数
	KeyPostVotedZSetPrefix = "post:voted:" // zset: 记录用户投票类型，参数是postId
)

func getRedisKey(key string) string {
	return Prefix + key
}
