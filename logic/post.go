package logic

import (
	"github.com/iamleizz/bluebell/dao/mysql"
	"github.com/iamleizz/bluebell/models"
	"github.com/iamleizz/bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// snowflake 生成 id
	p.ID = snowflake.GenID()
	// 访问 mysql 插入数据
	return mysql.CreatePost(p)
}