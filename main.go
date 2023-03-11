package main

import (
	sum "exemple/hello/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/sum", sum.Sum)
	router.Run()
}
