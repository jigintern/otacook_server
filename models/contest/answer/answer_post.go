package answer

import(
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Model is DB default data
type Model struct {
	AnswerID	int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Answer is answers  data
type Answer struct {
	Model
	QuestionID int `gorm:"not null"`
	UserID int `gorm:"not null"`
	CookingName string `gorm:"not null"`
	CookingOutline string `gorm:"type:text;not null"`
}

func InsertAns(qi int, ui int, cn string, co string) {
	db := GetDBConn()

	// テーブルの作成
	db.AutoMigrate(&Answer{})

	// insert into answers
	name := Answer { QuestionID: qi, UserID: ui, CookingName: cn, CookingOutline: co }
	db.Create(&name)
}

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
	USER := "root"
	PASS := "root"
	PROTCOL := "tcp(db:3306)"
	DBNAME := "otacook"
	OPTION := "charset=utf8&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTCOL + "/" + DBNAME + "?" + OPTION

	return DBMS, CONNECT
}

