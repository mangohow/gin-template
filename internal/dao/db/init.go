package db

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mangohow/gin-template/conf"
)

var sqlDb *sqlx.DB

func InitMysql() error {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	var err error
	db, err := sqlx.ConnectContext(timeoutCtx, "mysql", conf.Config().Mysql.DataSourceName)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(int(conf.Config().Mysql.MaxOpenConns))
	db.SetMaxIdleConns(int(conf.Config().Mysql.MaxIdleConns))

	sqlDb = db

	return nil
}

func CloseMysql() {
	if sqlDb != nil {
		_ = sqlDb.Close()
	}
}

func DB() *sqlx.DB {
	return sqlDb
}
