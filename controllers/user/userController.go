package user

import(
	"github.com/gin-gonic/gin"
	"m/models/user"
	"net/http"
	"strconv"
)

func PostUser(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	un := ctx.PostForm("UserName")
	um := ctx.PostForm("UserMailaddress")
	up := ctx.PostForm("UserPassword")
	ur, _ := strconv.Atoi(ctx.PostForm("UserRated"))
	user.InsertUser(un, um, up, ur)
	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}