package redis

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePreVote = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

/*
	投票功能
*/
func PostVote(ctx context.Context, userID, postID string, value float64) (err error){
	// pipeline := rdb.Pipeline()
	postTime := rdb.ZScore(ctx, getKey(KeyPostTime), postID).Val()
	// 如果时间超过一周，则不能投票
	if float64(time.Now().Unix()) - postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 从 zset 中获取当前用户的投票记录
	ov := rdb.ZScore(ctx, getKey(KeyPostVotedPf+postID), userID).Val()
	var dir float64
	if value > ov {
		// 如果现在的值大于原来的值，则代表需要加上一票
		dir = 1
	} else {
		// 否则需要减掉一票哦
		dir = -1
	}
	// 计算差值，这里有多种情况
	// 如果用户现在点了赞成，且原本为反对，或者点了反对，原本为赞成，则差值为 2
	// 其他情况差值为 1
	diff := math.Abs(ov - value)
	// 在数据库中加上对应的数值
	// 差值倍率 * (正 or 负) * 每一票代表的分值
	_, err = rdb.ZIncrBy(ctx, getKey(KeyPostScore), diff * dir * scorePreVote, postID).Result()
	
	if err != nil {
		return err
	}
	// 如果当前值为0
	if value == 0 {
		_, err = rdb.ZRem(ctx, getKey(KeyPostVotedPf + postID), postID).Result()
	} else {
		_, err = rdb.ZAdd(ctx, getKey(KeyPostVotedPf + postID), redis.Z{
			Member: userID,
			Score: value,
		}).Result()
	}

	return 
}

func CreatePost(ctx context.Context, postID int64) error {
	// 创建帖子时，将帖子的创建时间加入到 zset 中
	_, err := rdb.ZAdd(ctx, getKey(KeyPostTime), redis.Z{
		Score: float64(time.Now().Unix()),
		Member: postID,
	}).Result()
	return err
}