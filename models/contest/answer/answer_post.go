package answer

import(
	"time"
	"m/models"
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
	db := model.GetDBConn()

	// テーブルの作成
	db.AutoMigrate(&Answer{})

	// insert into answers
	name := Answer { QuestionID: qi, UserID: ui, CookingName: cn, CookingOutline: co }
	db.Create(&name)
}
