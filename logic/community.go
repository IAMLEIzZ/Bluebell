package logic

import (
	"github.com/iamleizz/bluebell/dao/mysql"
	"github.com/iamleizz/bluebell/models"
)

func GetCommunityList() (data []*models.CommunityList, err error) {
	return mysql.GetCommunityList()
}