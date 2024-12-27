package logic

import (
	"context"
	"strconv"

	"github.com/iamleizz/bluebell/dao/redis"
	"github.com/iamleizz/bluebell/models"
	"go.uber.org/zap"
)


func PostVote(ctx context.Context, userID int64, p *models.ParamVote) (err error){
	// 调用 dao 层的投票功能
	zap.L().Debug("PostVote", 
				zap.Int64("userID", userID),
				zap.String("postID", p.PostID),
				zap.Int8("direction", p.Direction))	
	return redis.PostVote(ctx, strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}