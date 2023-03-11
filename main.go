package main

import (
	"exemple/hello/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/sum", controller.Sum)
	router.Run()
}
