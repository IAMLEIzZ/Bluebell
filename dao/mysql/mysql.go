package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iamleizz/bluebell/setting"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)


var db *sqlx.DB

func Init(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Db,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))		
		return
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return 
}

func Close() {
	_ = db.Close()
}