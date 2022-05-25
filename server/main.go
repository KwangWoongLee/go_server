package main

import (
	"fmt"

	"github.com/KwangWoongLee/go_webserver/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("hi~")

	router := gin.Default()
	router.GET("/", handler.TestHandlerFunc)
	router.Run(":80")
}
