package mysql

import "github.com/iamleizz/bluebell/models"

// CreatePost  创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id, title, summary, author_id, community_id)
	values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Summary, p.AuthorID, p.CommunityID)
	return 
}

// GetPostDetail  获取帖子详细内容
func GetPostDetail(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, summary, author_id, community_id, create_time
			from post
			where post_id = ?`
	err = db.Get(post, sqlStr, pid)

	return 
}

// GetPostList  获取帖子列表
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title, summary, author_id, community_id, create_time
			from post
			limit ?, ?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page - 1) * size, size)
	return 
}