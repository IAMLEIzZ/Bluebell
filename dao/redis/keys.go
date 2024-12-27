package redis

const (
	KeyPrefix = "bluebell"
	KeyPostTime = "post:time"		// 记录帖子发帖时间
	KeyPostScore = "post:score"		// 记录帖子投票分数
	KeyPostVotedPf = "post:voted"	// 记录用户和投票类型
) 

func getKey(key string) string {
	return KeyPrefix + key
}