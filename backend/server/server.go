package server

import (
	"e-wallet/handler"
	"e-wallet/middleware"

	"github.com/gin-gonic/gin"
)

var (
	pathUsers         = "/users"
	pathUserPicture   = "/users/pictures"
	pathLogin         = "/users/login"
	pathTransfers     = "/transactions/transfer"
	pathTopup         = "/transactions/topup"
	pathTransactions  = "/transactions"
	pathResetPassword = "/reset-password"
	pathGameGacha     = "/games/gacha"
)

func Setup(opt RouterOpt) *gin.Engine {
	r := gin.Default()

	r.Static("/public/profile_pictures/", "profile_pictures/")

	r.ContextWithFallback = true

	r.Use(middleware.RequestId)
	r.Use(middleware.Logger(opt.Logger))
	r.Use(middleware.CustomMiddlewareError)
	r.Use(middleware.CorsMiddleware())

	r.POST(pathLogin, opt.UserHandler.LoginUser)

	r.POST(pathUsers, opt.UserHandler.RegisterUser)
	r.GET(pathUsers, middleware.JwtAuthMiddleware(), opt.UserHandler.GetDetail)
	r.PATCH(pathUsers, middleware.JwtAuthMiddleware(), opt.UserHandler.UpdateUserData)

	r.PATCH(pathUserPicture, middleware.JwtAuthMiddleware(), opt.UserHandler.UpdateProfilePicture)

	r.POST(pathTransfers, middleware.JwtAuthMiddleware(), opt.WalletHandler.Transfer)

	r.POST(pathTopup, middleware.JwtAuthMiddleware(), opt.WalletHandler.Topup)

	r.GET(pathTransactions, middleware.JwtAuthMiddleware(), opt.TransactionHandler.GetTransactionList)

	r.GET(pathResetPassword, opt.RPTHandler.RequestToken)
	r.POST(pathResetPassword, opt.RPTHandler.ResetPassword)

	r.POST(pathGameGacha, middleware.JwtAuthMiddleware(), opt.GameHandler.AttemptGacha)

	r.NoRoute(func(ctx *gin.Context) {
		handler.NoRouteHandlerFunc(ctx)
	})

	return r
}
