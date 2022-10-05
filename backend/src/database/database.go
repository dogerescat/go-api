package database

import (
	"database/sql"
	"fmt"
	"go-api/config"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func init() {
	connect()
}

func connect() {
	dbconf := config.Config.User + ":" + config.Config.Password + "@tcp(" + config.Config.Host + ":" + config.Config.Port + ")/" + config.Config.Database + "?charset=utf8mb4"
	Db, err = sql.Open("mysql", dbconf)
	if err != nil {
		fmt.Println(err.Error())
	}
}
