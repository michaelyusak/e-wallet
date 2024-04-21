package service

import (
	"context"
	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/entity"
	"e-wallet/helper"
	"e-wallet/repository"
)

type TransactionService interface {
	GetTransactionList(ctx context.Context, paginationParam entity.PaginationParameter) (*entity.TransactionList, error)
}

type transactionServiceIpl struct {
	transactionRepository repository.TransactionRepository
	walletRepository      repository.WalletRepository
}

func NewTransactionServiceIpl(transactionRepository repository.TransactionRepository, walletRepository repository.WalletRepository) transactionServiceIpl {
	return transactionServiceIpl{
		transactionRepository: transactionRepository,
		walletRepository:      walletRepository,
	}
}

func (s *transactionServiceIpl) GetTransactionList(ctx context.Context, paginationParam entity.PaginationParameter) (*entity.TransactionList, error) {
	err := helper.CheckPaginationParams(&paginationParam)
	if err != nil {
		return nil, err
	}

	id := ctx.Value(constants.UserId).(int)
	if id == 0 {
		return nil, apperror.StatusUnauthorized()
	}

	wallet, err := s.walletRepository.GetWalletByUserId(ctx, id)
	if err != nil {
		return nil, apperror.InternalServerErr(err.Error())
	}
	if wallet == nil {
		wallet = &entity.Wallet{}
	}

	transactionList, err := s.transactionRepository.GetTransactions(ctx, wallet.Number, paginationParam)
	if err != nil {
		return nil, apperror.InternalServerErr("internal error while fetching transaction list")
	}

	helper.FormatTransactionList(transactionList, wallet.Number, paginationParam.Page)

	return transactionList, nil
}
