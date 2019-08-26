package main

import(
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	Id	int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Answer struct {
	Model
	Cooking_name string
}

func main() {
	db := GetDBConn()

	// テーブルの作成
	db.AutoMigrate(&Answer{})

	// insert into answers
	name := Answer{Cooking_name: "肉じゃが"}
	db.Create(&name)
}

func GetDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	return db
}

func GetDBConfig() (string, string) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTCOL := "tcp(db:3306)"
	DBNAME := "otacook"
	OPTION := "charset=utf8&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTCOL + "/" + DBNAME + "?" + OPTION

	return DBMS, CONNECT
}

