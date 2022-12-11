package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"math"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一次投票的分数
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeat     = errors.New("你已经投过票了")
)

/*
投票的几种情况：
direction=1，有两种情况：
	1.之前没有投过票，现在投赞成票
	2.之前投反对票，现在改投赞成票
direction=0时，有两种情况：
	1.之前投过赞成票，现在要取消投票
	2.之前投过反对票，现在要取消投票
direction=-1时，有两种情况：
	1.之前没有投过票，现在投反对票
	2.之前投过赞成票，现在改投反对票

投票限制：
每个帖子发表之日起一个星期之内允许用户投票，超过一个星期就不允许投票了
	1.到期之后将redis中保存的赞成票数及反对票数存储到mysql中
	2.到期之后删除那个 KeyPostVotedZSetPF
*/

func VoteForPost(userId, postId string, value float64) error {
	// 1.判断投票限制
	postTime := client.ZScore(context.Background(), getRedisKey(KeyPostTimeZSet), postId).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2和3需要放到一个pipeline事务中操作
	pipeline := client.TxPipeline()

	// 2.更新帖子分数
	ov := client.ZScore(context.Background(), getRedisKey(KeyPostVotedZSetPrefix+postId), userId).Val()
	if value == ov {
		return ErrVoteRepeat
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	pipeline.ZIncrBy(context.Background(), getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postId)

	// 3.记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(context.Background(), getRedisKey(KeyPostVotedZSetPrefix+postId), postId)
	} else {
		pipeline.ZAdd(context.Background(), getRedisKey(KeyPostVotedZSetPrefix+postId), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userId,
		})
	}
	_, err := pipeline.Exec(context.Background())
	return err
}

func CreatePost(postId int64) error {
	pipeline := client.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	// 帖子分数
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	_, err := pipeline.Exec(context.Background())
	return err
}
