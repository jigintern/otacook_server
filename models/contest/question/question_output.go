package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// dbにアクセス
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/questionsdb ")
	// エラーハンドリング
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// SQLを実行
	rows, err := db.Query("SELECT * FROM question")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	// 結果を変数に代入し、出力する
	for rows.Next() {
		var question_id  int
		var question_name string
		var question_text string

		if err := rows.Scan(&question_id , &question_name , &question_text); err != nil {
			panic(err.Error())
		}
		fmt.Println(question_id , question_name, &question_text)
	}
}