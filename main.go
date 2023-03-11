package main

import (
	"exemple/hello/src/controller/subtract"
	"exemple/hello/src/controller/sum"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/sum", sum.Sum)
	router.GET("/subtract", subtract.Subtract)
	router.Run()
}
