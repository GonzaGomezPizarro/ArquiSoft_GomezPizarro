package main

import (
	"github.com/gin-gonic/gin" // tira error
)

func main() {
	engine := gin.New()
	engine.GET("/ping", Ping)
	engine.Run(":3000")

}
