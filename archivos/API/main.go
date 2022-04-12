package main

import (
	"github.com/gin-gogo gnic/gin" // tira error
)

func main() {
	engine := gin.New()
	engine.GET("/ping", Ping)
	engine.Run(":3000")

}

func Test (ctx *gin.Context) {
	var body body
	err:= ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, body)
}

func Ping (ctx *gin.Context) {
	ctx.String(200, "pong")
}