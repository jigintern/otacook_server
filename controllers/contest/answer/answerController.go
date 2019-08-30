package answer

import(
	"github.com/gin-gonic/gin"
	"m/models/contest/answer"
	"net/http"
	"strconv"
)

func PostAnswer(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	qi, _ := strconv.Atoi(ctx.PostForm("qi"))
	ui, _ := strconv.Atoi(ctx.PostForm("ui"))
	cn := ctx.PostForm("cn")
	co := ctx.PostForm("co")

	answer.InsertAns(qi, ui, cn, co)
	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}