package mysql

import (
	"database/sql"

	"github.com/iamleizz/bluebell/models"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.CommunityList, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)

	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no commyinty in db")
			err = nil
		}
	}

	return
}

func GetCommunityDetail(id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	sqlStr := "select community_id, community_name, introduction, create_time from community where community_id = ?"
	err = db.Get(communityDetail, sqlStr, id)
	
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrInvalidID	
		}
	}
	return communityDetail, err
}