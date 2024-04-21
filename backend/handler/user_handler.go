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

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterUser(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var req dto.UserDTO

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	userReq := dto.ToUser(req)

	user, token, err := h.userService.RegisterUser(ctx.Request.Context(), userReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.MessageResponse{
		Message: constants.MsgResUserCreated,
		Data: gin.H{
			"user":  dto.ToUserDTO(*user),
			"token": token,
		},
	})
}

func (h *UserHandler) LoginUser(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var req dto.UserDTO

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	userReq := dto.ToUser(req)

	token, err := h.userService.Login(ctx.Request.Context(), userReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageResponse{
		Message: constants.MsgResLoginSuccess,
		Data: dto.TokenJwtDTO{
			Token: token,
		},
	})
}

func (h *UserHandler) UpdateUserData(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	id, isExist := ctx.Get(string(constants.UserId))
	if !isExist {
		ctx.Error(apperror.StatusUnauthorized())
		return
	}

	var req dto.UserDataDTO

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	c := context.WithValue(ctx, constants.UserId, id)

	userReq := dto.ToUserData(req)

	err = h.userService.UpdateUserData(c, userReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageResponse{
		Message: "user data updated",
	})
}

func (h *UserHandler) GetDetail(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	id, isExist := ctx.Get(string(constants.UserId))
	if !isExist {
		ctx.Error(apperror.StatusUnauthorized())
		return
	}

	c := context.WithValue(ctx, constants.UserId, id)

	user, err := h.userService.GetDetail(c)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageResponse{
		Message: constants.MsgResAccFound,
		Data:    dto.ToUserDTO(*user),
	})
}

func (h *UserHandler) UpdateProfilePicture(ctx *gin.Context) {
	userId, isExist := ctx.Get(string(constants.UserId))
	if !isExist {
		ctx.Error(apperror.StatusUnauthorized())
		return
	}

	uploadedImage, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.Error(err)
		return
	}

	c := context.WithValue(ctx, constants.UserId, userId)

	err = h.userService.UpdateProfilePicture(c, uploadedImage, fileHeader)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.MessageResponse{
		Message: "profile picture updated!",
	})
}
