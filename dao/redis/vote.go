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

func PostVote(ctx context.Context, userID, postID string, value float64) (err error){
	// pipeline := rdb.Pipeline()
	postTime := rdb.ZScore(ctx, getKey(KeyPostTime), postID).Val()
	x := float64(time.Now().Unix()) - postTime
	if x > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	ov := rdb.ZScore(ctx, getKey(KeyPostVotedPf+postID), userID).Val()
	var dir float64
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value)
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
	_, err := rdb.ZAdd(ctx, getKey(KeyPostTime), redis.Z{
		Score: float64(time.Now().Unix()),
		Member: postID,
	}).Result()
	return err
}