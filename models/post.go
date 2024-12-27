package models

import "time"

// 帖子内容信息
type Post struct {
	ID          int64     `json:"post_id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id,string" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Summary     string    `json:"summary" db:"summary" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// 帖子详细信息包含作者名称
type PostDetail struct {
	AuthorName       string `json:"author_name"`
	*Post            `json:"post"`
	*CommunityDetail `json:"community"`
}
