package sum

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Sum(gin *gin.Context) {
	gin.JSON(http.StatusOK, "hello")
}
