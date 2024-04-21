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

type WalletHandler struct {
	walletService service.WalletService
}

func NewWalletHandler(walletService service.WalletService) WalletHandler {
	return WalletHandler{
		walletService: walletService,
	}
}

func (h *WalletHandler) Transfer(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var req dto.TransactionDTO

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	id, isExist := ctx.Get(string(constants.UserId))
	if !isExist {
		ctx.Error(apperror.StatusUnauthorized())
		return
	}

	c := context.WithValue(ctx, constants.UserId, id)

	transferReq := dto.ToTransaction(req)

	err = h.walletService.Transfer(c, transferReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageResponse{
		Message: constants.MsgResTransferSuccess,
	})
}

func (h *WalletHandler) Topup(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var req dto.TransactionDTO

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	id, isExist := ctx.Get(string(constants.UserId))
	if !isExist {
		ctx.Error(apperror.StatusUnauthorized())
		return
	}

	c := context.WithValue(ctx, constants.UserId, id)

	topupReq := dto.ToTransaction(req)

	err = h.walletService.Topup(c, topupReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageResponse{
		Message: constants.MsgResTopupSuccess,
	})
}
