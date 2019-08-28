package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"m/routes"
	//"m/models/contest/answer"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	router.POST("/answer/insert",routes.UserAnswer)
	router.GET("/", routes.Home)
	router.GET("/answer", routes.Answer)

	user := router.Group("/user")
	{
		user.POST("/signup", routes.UserSignUp)
		user.POST("/login", routes.UserLogIn)
	}

	router.GET("/login", routes.LogIn)
	router.GET("/signup", routes.SignUp)
	router.NoRoute(routes.NoRoute)

	router.Run(":8080")

	/*
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

	answer.InsertAns(qi, ui, cn, co)
	*/
}