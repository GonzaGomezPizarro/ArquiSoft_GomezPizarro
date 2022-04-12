package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // tira error
)

func main() {
	engine := gin.New()
	engine.GET("/ping", Ping)
	engine.Run(":3000")

}

type Body struct {
	Name string `json:"name"`
}

func Test(ctx *gin.Context) {
	var body Body
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, body)
}

func Ping(ctx *gin.Context) {
	ctx.String(200, "pong")
}
