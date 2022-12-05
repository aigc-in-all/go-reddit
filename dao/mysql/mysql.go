package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"goreddit/setting"
)

var db *sqlx.DB

func Init(cfg *setting.MySQLConfig) (err error) {
	// user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenCons)
	db.SetMaxIdleConns(cfg.MaxIdleCons)
	return
}

func Close() {
	_ = db.Close()
}
