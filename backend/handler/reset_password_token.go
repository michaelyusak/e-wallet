package handler

import (
	"net/http"

	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/dto"
	"e-wallet/service"

	"github.com/gin-gonic/gin"
)

type RPTHandler struct {
	rptService service.RPTService
}

func NewRPTHandler(rptService service.RPTService) RPTHandler {
	return RPTHandler{
		rptService: rptService,
	}
}

func (h *RPTHandler) RequestToken(ctx *gin.Context) {
	var req dto.ResetPasswordTokenDTO

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	rptReq := dto.ToResetPasswordToken(req)

	rpt, err := h.rptService.RequestToken(ctx.Request.Context(), rptReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageResponse{
		Message: constants.MsgResOK,
		Data:    dto.ToResetPasswordTokenDTO(*rpt),
	})
}

func (h *RPTHandler) ResetPassword(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var req dto.ResetPasswordTokenDTO

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(apperror.InternalServerErr(err.Error()))
		return
	}

	resetPwdReq := dto.ToResetPasswordToken(req)

	err = h.rptService.ResetPassword(ctx.Request.Context(), resetPwdReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageResponse{
		Message: constants.MsgResOK,
	})
}
