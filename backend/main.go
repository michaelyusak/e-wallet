package main

import (
	"context"
	"e-wallet/config"
	"e-wallet/database"
	"e-wallet/handler"
	"e-wallet/helper"
	"e-wallet/repository"
	"e-wallet/server"
	"e-wallet/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := config.ConfigInit(); err != nil {
		log.Fatal("error while loading env")
	}

	addr := os.Getenv("ADDR")

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("error while connecting to db")
	}

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	userRepository := repository.NewUserRepositoryPostgres(db)
	walletRepository := repository.NewWalletRepositoryPostgres(db)
	transactionRepository := repository.NewTransactionRepositoryPostgres(db)
	rptRepository := repository.NewRPTRepositoryPostgres(db)
	prizeRepository := repository.NewPrizeRepositoryPostgres(db)
	gachaRepository := repository.NewGachaRepositoryPostgres(db)
	hashHelper := new(helper.HashHelperImpl)
	tokenHelper := new(helper.TokenHelperImpl)

	userService := service.NewUserServiceIpl(&userRepository, &walletRepository, hashHelper, tokenHelper)
	walletService := service.NewWalletServiceIpl(&walletRepository, &userRepository, &transactionRepository)
	transactionService := service.NewTransactionServiceIpl(&transactionRepository, &walletRepository)
	rptService := service.NewRPTService(&rptRepository, &userRepository, tokenHelper, hashHelper)
	gameService := service.NewGameServiceIpl(&prizeRepository, &gachaRepository, &walletRepository)

	UserHandler := handler.NewUserHandler(&userService)
	walletHandler := handler.NewWalletHandler(&walletService)
	transactionHandler := handler.NewTransactionHandler(&transactionService)
	rptHandler := handler.NewRPTHandler(&rptService)
	gameHandler := handler.NewGameHandler(&gameService)

	router := server.Setup(server.RouterOpt{
		UserHandler:        UserHandler,
		WalletHandler:      walletHandler,
		TransactionHandler: transactionHandler,
		RPTHandler:         rptHandler,
		GameHandler:        gameHandler,
		Logger:             log,
	})

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	<-ctx.Done()
}
