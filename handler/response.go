package handler

import "github.com/gin-gonic/gin"

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSucess(ctx *gin.Context, op string,data interface{}) {

}
