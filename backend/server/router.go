package server

import (
	"e-wallet/handler"

	"github.com/sirupsen/logrus"
)

type RouterOpt struct {
	UserHandler        handler.UserHandler
	WalletHandler      handler.WalletHandler
	TransactionHandler handler.TransactionHandler
	RPTHandler         handler.RPTHandler
	GameHandler        handler.GameHandler
	Logger             *logrus.Logger
}
