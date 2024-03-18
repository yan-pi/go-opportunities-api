package handler

import (
	"fmt"
	"net/http"
	"github.com/yan-pi/go-opportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponde `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponde `json:"data"`
}
type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponde `json:"data"`
}
type ListOpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponde `json:"data"`
}
type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponde `json:"data"`
}
