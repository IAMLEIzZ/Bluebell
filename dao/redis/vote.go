package redis

import (
	"context"
	"math"
	"time"

	"github.com/redis/go-redis/v9"
)

/*
	投票功能
*/
func PostVote(ctx context.Context, userID, postID string, value float64) (err error){
	// pipeline := rdb.Pipeline()
	postTime := rdb.ZScore(ctx, getKey(KeyPostTime), postID).Val()
	// 如果时间超过一周，则不能投票
	if float64(time.Now().Unix()) - postTime > OneWeekInSeconds {
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
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(ctx, getKey(KeyPostScore), diff * dir * ScorePreVote, postID)
	// 如果当前值为0
	if value == 0 {
		pipeline.ZRem(ctx, getKey(KeyPostVotedPf + postID), postID)
	} else {
		pipeline.ZAdd(ctx, getKey(KeyPostVotedPf + postID), redis.Z{
			Member: userID,
			Score: value,
		})
	}
	_, err = pipeline.Exec(ctx)

	return 
}

func CreatePost(ctx context.Context, postID int64) (err error) {
	// redis 事务
	pipeline := rdb.TxPipeline()
	// 创建帖子时，将帖子的创建时间加入到 zset 中
	pipeline.ZAdd(ctx, getKey(KeyPostTime), redis.Z{
		Score: float64(time.Now().Unix()),
		Member: postID,
	})

	// 创建帖子时，将帖子的默认分值加入 zset 中
	pipeline.ZAdd(ctx, getKey(KeyPostScore), redis.Z{
		Score: float64(time.Now().Unix()),
		Member: postID,
	})

	_, err = pipeline.Exec(ctx)

	return err
}