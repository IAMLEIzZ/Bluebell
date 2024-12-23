package logic

import (
	"github.com/iamleizz/bluebell/dao/mysql"
	"github.com/iamleizz/bluebell/models"
	"github.com/iamleizz/bluebell/pkg/jwt"
	"github.com/iamleizz/bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUP) (err error){

	// 判断用户是否存在
	if err := mysql.CheckUserExists(p.Username); err != nil {
		return err
	}
	userID := snowflake.GenID()
	user := &models.User{
		UserID: userID,
		Username: p.Username,
		Password: p.Password,
	}	
	
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	
	if err := mysql.Login(user); err != nil {
		return "", err
	}

	return jwt.GenToken(user.UserID, user.Username)
}