package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"goreddit/model"
)

func GetPostIdsInOrder(p *model.ParamPostList) ([]string, error) {
	// 1.根据参数确定Key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == model.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2.确定查询的起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// 3.ZRevRange 按分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(context.Background(), key, start, end).Result()
}

// GetPostVoteData 根据ids查询每篇帖子的投赞成票数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	/*data = make([]int64, 0, len(ids))
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		// 查找 key 中分数是1的元素的数量->统计每篇帖子的赞成票的数量
		v := client.ZCount(context.Background(), key, "1", "1").Val()
		data = append(data, v)
	}*/
	pipline := client.TxPipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipline.ZCount(context.Background(), key, "1", "1")
	}
	cmders, err := pipline.Exec(context.Background())
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(ids))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return data, nil
}
