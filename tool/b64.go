package tool

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
)

type text struct {
	Content string`json:"content"`
}

func B64En(ctx *gin.Context)  {
	var json text
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"message": err.Error(),
		})
		return
	}
	result := base64.StdEncoding.EncodeToString([]byte(json.Content))
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": result,
		"message": "",
	})
}

func B64De(ctx *gin.Context)  {
	var json text
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"message": err.Error(),
		})
		return
	}
	result, err := base64.StdEncoding.DecodeString(json.Content)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": string(result),
		"message": "",
	})
}
