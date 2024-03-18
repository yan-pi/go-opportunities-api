package handler

import (
	"net/http"

	"github.com/Yan-pi/go-opportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error list openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
