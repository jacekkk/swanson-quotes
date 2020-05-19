package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	cfg := mysql.Cfg("swanson-quotes-277514:europe-west2:swanson-quotes", "root", "9n15vPsrkH6Kxrua")
	cfg.DBName = "swanson_quotes"
	db, err = mysql.DialCfg(cfg)

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	fmt.Println("Connected to DB")
}
