package redis

import (
	"context"

	"github.com/iamleizz/bluebell/models"
)

func GetPostIdsByOrder(ctx context.Context, p *models.ParamPostList) ([]string, error) {
	// 获取 Key
	key := getKey(KeyPostTime)
	if p.Order == PostOrderByScore {
		key = getKey(KeyPostScore)
	}
	// 计算其实起始参数
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	return rdb.ZRevRange(ctx, key, start, end).Result()
}