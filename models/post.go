package models

import "time"

type Post struct {
	ID int64 `json:"post_id" db:"post_id"`
	AuthorID int64 `json:"author_id" db:"author_id"`
	CommunityID int64 `json:"community_id" db:"community_id" binding:"required"`
	Status int32 `json:"status" db:"status"`
	Title string `json:"title" db:"title" binding:"required"`
	Summary string `json:"summary" db:"summary" binding:"required"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}