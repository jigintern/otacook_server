package model

import	(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetDBConn is the function that connects to the DB
func GetDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	return db
}

// GetDBConfig is the function that enter the DB config.
func GetDBConfig() (string, string) {
	DBMS := "mysql"
	USER := "testuser1"
	PASS := "pass1"
	PROTCOL := "tcp(db:3306)"
	DBNAME := "otacook"
	OPTION := "charset=utf8&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTCOL + "/" + DBNAME + "?" + OPTION

	return DBMS, CONNECT
}


