package redis

import "errors"

const (
	PostOrderByTime = "time"
	PostOrderByScore = "score"
)

const (
	OneWeekInSeconds = 7 * 24 * 3600
	ScorePreVote = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)