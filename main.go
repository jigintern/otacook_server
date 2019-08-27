package main

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"m/models/contest/answer"
)

func main() {
	var qi, ui int
	var cn, co string
	fmt.Print("qi=")
	fmt.Scan(&qi)
	fmt.Print("ui=")
	fmt.Scan(&ui)
	fmt.Print("cn=")
	fmt.Scan(&cn)
	fmt.Print("co=")
	fmt.Scan(&co)

	//router := gin.Default()
	//router.LoadHTMLGlob("views/*.html")
	//routar.Static("/assets", "./assets")
	//router.GET("/", func(c *gin.Context) {
		//c.HTML()
//	})
//)

	answer.InsertAns(qi, ui, cn, co)
}
