package routes

import (
	"github.com/gin-gonic/gin"
	"m/models/contest/answer"
	"net/http"
	"strconv"
)

func UserAnswer(ctx *gin.Context) {
	qi, _ := strconv.Atoi(ctx.PostForm("questionid"))
	ui, _ := strconv.Atoi(ctx.PostForm("userid"))
	cn := ctx.PostForm("cookingname")
	co := ctx.PostForm("cookingoutline")

	println("Answer_post")
	println(qi)
	println(ui)
	println(cn)
	println(co)


	answer.InsertAns(qi, ui, cn, co)
	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}



func UserSignUp(ctx *gin.Context) {
	println("post/signup")
	username := ctx.PostForm("username")
	email := ctx.PostForm("emailaddress")
	password := ctx.PostForm("password")
	passwordConf := ctx.PostForm("passwordconfirmation")
	println("username: " + username)
	println("email: " + email)
	println("password: " + password)
	println("passwordConf: " + passwordConf)

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}

func UserLogIn(ctx *gin.Context) {
	println("post/login")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	println("username: " + username)
	println("password: " + password)

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}
