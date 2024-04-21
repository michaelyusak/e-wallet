package middleware

import (
	"net/http"

	"e-wallet/constants"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId(ctx *gin.Context) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Set(string(constants.RequestId), uuid)
	ctx.Next()
}
