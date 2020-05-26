package models

import (
	"database/sql"
	"fmt"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(user string, password string, connectionName string, dbName string, env string) {
	var err error

	if env == "PROD" {
		dbURI := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", user, password, connectionName, dbName)
		db, err = sql.Open("mysql", dbURI)
	} else {
		cfg := mysql.Cfg(connectionName, user, password)
		cfg.DBName = dbName
		db, err = mysql.DialCfg(cfg)
	}

	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}

	fmt.Println("db connected")
}
