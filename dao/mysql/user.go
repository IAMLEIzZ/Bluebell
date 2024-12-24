package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	"github.com/iamleizz/bluebell/models"
)

const secret string = "IAMLEIzZ"

func CheckUserExists(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrUserExist
	}
	return nil
}

func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)

	return err
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	if err = db.Get(user, sqlStr, user.Username); err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotExist
		}
		return err
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrInvalidPassword
	}
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}