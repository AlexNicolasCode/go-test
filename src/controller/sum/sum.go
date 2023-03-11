package sum

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Sum(g *gin.Context) {
	firstnumber, _ := strconv.Atoi(g.Query("firstnumber"))
	secondnumber, _ := strconv.Atoi(g.Query("secondnumber"))
	g.JSON(http.StatusOK, secondnumber+firstnumber)
}
