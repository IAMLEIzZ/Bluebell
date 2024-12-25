package logic

import (
	"github.com/iamleizz/bluebell/dao/mysql"
	"github.com/iamleizz/bluebell/models"
	"github.com/iamleizz/bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// snowflake 生成 id
	p.ID = snowflake.GenID()
	// 访问 mysql 插入数据
	return mysql.CreatePost(p)
}

func GetPostDetail(pid int64) (PostDetail *models.PostDetail, err error) {
	post, err := mysql.GetPostDetail(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostDetail failed", zap.Error(err))
		return nil, err
	}

	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID failed", zap.Error(err))
		return nil, err
	}
	commuinty, err := mysql.GetCommunityDetail(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetail failed", zap.Error(err))
		return nil, err
	}
	
	PostDetail = &models.PostDetail{
		AuthorName: user.Username,
		Post: post,
		CommunityDetail: commuinty,
	}
	return 
}