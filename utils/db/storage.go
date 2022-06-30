package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/url"
	"time"
)

var (
	client *MysqlClient
)

func Init(cfg DBConfig) error {
	var err error
	client, err = NewMysqlClient(
		getDSN(cfg),
		60,
		30,
		5*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func getDSN(conf DBConfig) string {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=true&loc=Local&time_zone=%v",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Database, url.QueryEscape("'Asia/Shanghai'"))
	return dsn
}

func GetDb() *sqlx.DB {
	return client.GetDb()
}
