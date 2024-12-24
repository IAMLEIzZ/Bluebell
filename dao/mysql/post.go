package mysql

import "github.com/iamleizz/bluebell/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id, title, summary, author_id, community_id)
	values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Summary, p.AuthorID, p.CommunityID)
	return 
}