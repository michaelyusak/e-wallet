package handler

import (
	"net/http"

	"e-wallet/dto"

	"github.com/gin-gonic/gin"
)

func NoRouteHandlerFunc(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, dto.ErrResponse{
		Message: "page not found",
		Code:    http.StatusNotFound,
	})
}
