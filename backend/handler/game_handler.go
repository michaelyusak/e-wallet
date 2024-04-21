package handler

import (
	"context"
	"net/http"

	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/dto"
	"e-wallet/service"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	gameService service.GameService
}

func NewGameHandler(gameService service.GameService) GameHandler {
	return GameHandler{
		gameService: gameService,
	}
}

func (h *GameHandler) AttemptGacha(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	id, isExist := ctx.Get(string(constants.UserId))
	if !isExist {
		ctx.Error(apperror.StatusUnauthorized())
		return
	}

	c := context.WithValue(ctx, constants.UserId, id)

	selectionStr := ctx.Query(constants.GameGachaSelection)

	selectedBox, err := h.gameService.AttemptGacha(c, selectionStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageResponse{
		Message: constants.MsgResOK,
		Data:    selectedBox,
	})
}
