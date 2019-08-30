package user

import(
	"time"
	"m/models"
)

// Model is DB default data
type Model struct {
	UserID	int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Answer is answers  data
type User struct {
	Model
	UserName string `gorm:"not null"`
	UserMailaddress string `gorm:"not null"`
	UserPassword string `gorm:"not null"`
	UserRated int `gorm:"not null"`
}

func InsertUser(un string, um string, up string, ur int) {
	db := model.GetDBConn()

	// テーブルの作成
	db.AutoMigrate(&User{})

	// insert into answers
	name := User{ UserName: un, UserMailaddress: um, UserPassword: up, UserRated: ur }
	db.Create(&name)
}