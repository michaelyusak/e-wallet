package service

import (
	"context"
	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/entity"
	"e-wallet/helper"
	"e-wallet/repository"
	"errors"

	"github.com/shopspring/decimal"
)

type WalletService interface {
	Topup(ctx context.Context, topupReq entity.Transaction) error
	Transfer(ctx context.Context, transferReq entity.Transaction) error
}

type walletServiceIpl struct {
	walletRepository      repository.WalletRepository
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewWalletServiceIpl(walletRepository repository.WalletRepository, userRepository repository.UserRepository, transactionRepository repository.TransactionRepository) walletServiceIpl {
	return walletServiceIpl{
		walletRepository:      walletRepository,
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

var (
	gachaTrialBaseAmount = decimal.NewFromInt(10000000)
)

func (s *walletServiceIpl) Transfer(ctx context.Context, transferReq entity.Transaction) error {
	senderId := ctx.Value(constants.UserId).(int)
	if senderId == 0 {
		return apperror.StatusUnauthorized()
	}

	walletS, err := s.walletRepository.GetWalletByUserId(ctx, senderId)
	if err != nil {
		return apperror.InternalServerErr("internal error while checking user wallet")
	}
	if walletS == nil {
		return apperror.StatusAccountUnregistered()
	}

	if walletS.Balance.LessThan(transferReq.Amount) {
		return apperror.StatusUnsufficientBalance()
	}

	transferReq.SenderWalletNum = walletS.Number

	err = helper.CheckForTransfer(&transferReq)
	if err != nil {
		return err
	}

	walletR, err := s.walletRepository.GetWalletByNum(ctx, transferReq.RecepientWalletNum)
	if err != nil {
		return apperror.InternalServerErr("internal error while checking recepient wallet")
	}
	if walletR == nil {
		return apperror.StatusRecepientNotFound()
	}

	errWithTx := s.walletRepository.WithTx(ctx, func(repo repository.WalletRepository) error {
		err = repo.UpdateBalance(ctx, walletR.Id, transferReq.Amount)
		if err != nil {
			return apperror.InternalServerErr("internal error while updating recepient balance")
		}

		err = repo.UpdateBalance(ctx, walletS.Id, transferReq.Amount.Neg())
		if err != nil {
			return apperror.InternalServerErr("internal error while updating sender balance")
		}

		err = repo.PostOneTransaction(ctx, transferReq)
		if err != nil {
			return apperror.InternalServerErr("internal error while posting transaction record")
		}

		return nil
	})

	if errWithTx != nil {
		if errors.As(errWithTx, &apperror.AppError{}) {
			return errWithTx
		}

		return apperror.InternalServerErr("internal error while executing transaction")
	}

	return nil
}

func (s *walletServiceIpl) Topup(ctx context.Context, topupReq entity.Transaction) error {
	recepientId := ctx.Value(constants.UserId).(int)
	if recepientId == 0 {
		return apperror.StatusUnauthorized()
	}

	walletS, err := helper.CheckForTopup(&topupReq)
	if err != nil {
		return err
	}

	walletR, err := s.walletRepository.GetWalletByUserId(ctx, recepientId)
	if err != nil {
		return apperror.InternalServerErr("internal error while checking recepient wallet")
	}
	if walletR == nil {
		return apperror.StatusRecepientNotFound()
	}

	topupReq.SenderWalletNum = walletS.Number
	topupReq.RecepientWalletNum = walletR.Number

	errWithTx := s.walletRepository.WithTx(ctx, func(repo repository.WalletRepository) error {
		err = repo.UpdateBalance(ctx, walletS.Id, topupReq.Amount.Neg())
		if err != nil {
			return apperror.InternalServerErr("internal error while contacting your source of fund")
		}

		err = repo.UpdateBalance(ctx, walletR.Id, topupReq.Amount)
		if err != nil {
			return apperror.InternalServerErr("internal error while updating recepient balance")
		}

		err = repo.PostOneTransaction(ctx, topupReq)
		if err != nil {
			return apperror.InternalServerErr("internal error while posting transaction record")
		}

		return nil
	})

	if errWithTx != nil {
		if errors.As(errWithTx, &apperror.AppError{}) {
			return errWithTx
		}

		return apperror.InternalServerErr("internal error while executing transaction")
	}

	if topupReq.Amount.GreaterThanOrEqual(gachaTrialBaseAmount) {
		getTrial := topupReq.Amount.Div(gachaTrialBaseAmount).Floor()

		err := s.walletRepository.UpdateGachaTrial(ctx, walletR.UserId, getTrial)
		if err != nil {
			return apperror.InternalServerErr("internal error while updating user gacha trial")
		}
	}

	return nil
}
