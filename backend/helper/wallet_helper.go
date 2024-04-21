package helper

import (
	"fmt"
	"strconv"

	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/entity"

	"github.com/shopspring/decimal"
)

var (
	minTfAmount = decimal.NewFromInt(1000)
	maxTfAmount = decimal.NewFromInt(50000000)

	minTpAmount = decimal.NewFromInt(50000)
	maxTpAmount = decimal.NewFromInt(10000000)
)

func CheckForTransfer(transferReq *entity.Transaction) error {
	if transferReq.RecepientWalletNum == "" {
		return apperror.BadRequest("should fill recepient wallet number to perform transfer")
	}

	if transferReq.Amount.LessThan(minTfAmount) || transferReq.Amount.GreaterThan(maxTfAmount) {
		return apperror.StatusAmountInvalid(minTfAmount, maxTfAmount)
	}

	transferReq.SourceOfFunds = constants.WalletStr

	if transferReq.SenderWalletNum == transferReq.RecepientWalletNum {
		return apperror.StatusSelfTransfer()
	}

	return nil
}

func CheckForTopup(topupReq *entity.Transaction) (*entity.Wallet, error) {
	var walletS entity.Wallet

	if topupReq.Amount.LessThan(minTpAmount) || topupReq.Amount.GreaterThan(maxTpAmount) {
		return &walletS, apperror.StatusAmountInvalid(minTpAmount, maxTpAmount)
	}

	if topupReq.SourceOfFunds == "" {
		return &walletS, apperror.StatusSourceOfFundsEmpty()
	}

	if topupReq.SourceOfFunds == constants.WalletStr {
		return &walletS, apperror.StatusRequestUnavailable()
	}

	if topupReq.SourceOfFunds == constants.GachaStr {
		return &walletS, apperror.StatusUnauthorized()
	}

	switch topupReq.SourceOfFunds {
	case constants.BankTransferStr:
		walletS.Id = constants.BankWalletId
		walletS.Number = fmt.Sprintf("100000000000%s", strconv.Itoa(constants.BankWalletId))

	case constants.CreditCardStr:
		walletS.Id = constants.BankWalletId
		walletS.Number = fmt.Sprintf("100000000000%s", strconv.Itoa(constants.BankWalletId))


	case constants.CashStr:
		walletS.Id = constants.RetailWalletId
		walletS.Number = fmt.Sprintf("100000000000%s", strconv.Itoa(constants.RetailWalletId))


	case constants.GachaStr:
		walletS.Id = constants.CompanyWalletId
		walletS.Number = fmt.Sprintf("100000000000%s", strconv.Itoa(constants.CompanyWalletId))
	}

	topupReq.Description = fmt.Sprintf("Top Up from %s", topupReq.SourceOfFunds)

	return &walletS, nil
}
