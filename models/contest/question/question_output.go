package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/questionsdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM question")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

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