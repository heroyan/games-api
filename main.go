package main

import (
	"github.com/heroyan/games-api/tool"

	"github.com/gin-gonic/gin"
)
func main()  {
	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/panic", fault)

	tools := r.Group("/api/tool")
	{
		tools.POST("/b64en", tool.B64En)
		tools.POST("/b64de", tool.B64De)
	}

	r.Run()
}

func ping(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func fault(ctx *gin.Context)  {
	panic("i am quiting")
}