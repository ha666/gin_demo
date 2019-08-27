package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ha666/gin_demo/initial/config"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var (
	ha666db *sqlx.DB
	err     error
)

func Init() {
	initHa666db()
}

func initHa666db() {
	ha666db, err = sqlx.Open("mysql", fmt.Sprintf(config.Config.DB.Dsn, config.Config.DB.Pwd))
	if err != nil {
		log.Fatalf("【initHa666db.NewEngine】ex:%s", err.Error())
		os.Exit(0)
		return
	}
	err = ha666db.Ping()
	if err != nil {
		log.Fatalf("【initHa666db.Ping】ex:%s", err.Error())
		os.Exit(0)
		return
	}
	ha666db.SetMaxIdleConns(config.Config.DB.MaxIdleConn)
	ha666db.SetMaxOpenConns(config.Config.DB.MaxOpenConn)
}
