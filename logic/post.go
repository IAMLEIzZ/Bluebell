package logic

import (
	"context"

	"github.com/iamleizz/bluebell/dao/mysql"
	"github.com/iamleizz/bluebell/dao/redis"
	"github.com/iamleizz/bluebell/models"
	"github.com/iamleizz/bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

// CreatePost  创建帖子
func CreatePost(ctx context.Context, p *models.Post) (err error) {
	// snowflake 生成 id
	p.ID = snowflake.GenID()
	// 访问 mysql 插入数据
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	return redis.CreatePost(ctx, p.ID)
}

// GetPostDetail  根据帖子 id 获取详细的帖子信息
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

// GetPostList  获取帖子列表
func GetPostList(page, size int64) (postlist []*models.PostDetail, err error) {
	// 这里 posts 的类型是 post 的切片
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	postlist = make([]*models.PostDetail, 0, len(posts))

	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID failed", zap.Error(err))
			continue
		}
		commuinty, err := mysql.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetail failed", zap.Error(err))
			continue
		}
		PostDetail := &models.PostDetail{
			AuthorName: user.Username,
			Post: post,
			CommunityDetail: commuinty,
		}
		postlist = append(postlist, PostDetail)
	}
	return
}