package dto

import (
	"e-wallet/constants"
	"e-wallet/entity"

	"github.com/shopspring/decimal"
)

type TransactionDTO struct {
	Id                 int             `json:"id,omitempty"`
	Type               string          `json:"type"`
	SenderWalletNum    string          `json:"from"`
	SenderName         string          `json:"sender_name"`
	RecepientWalletNum string          `json:"to"`
	RecepientName      string          `json:"recipient_name"`
	Amount             decimal.Decimal `json:"amount" binding:"required"`
	SourceOfFunds      int             `json:"source_of_funds"`
	Description        string          `json:"description" binding:"max=35"`
	Date               string          `json:"date,omitempty"`
}

type TransactionListDTO struct {
	TotalItem    int              `json:"total_item"`
	Page         string           `json:"page"`
	Transactions []TransactionDTO `json:"transactions"`
}

func ToTransaction(transactionDTO TransactionDTO) entity.Transaction {
	return entity.Transaction{
		Id:                 transactionDTO.Id,
		SenderWalletNum:    transactionDTO.SenderWalletNum,
		RecepientWalletNum: transactionDTO.RecepientWalletNum,
		Amount:             transactionDTO.Amount,
		SourceOfFunds:      sourceOfFundsToString(transactionDTO.SourceOfFunds),
		Description:        transactionDTO.Description,
	}
}

func sourceOfFundsToString(sof int) string {
	switch sof {
	case constants.Wallet:
		return constants.WalletStr
	case constants.BankTransfer:
		return constants.BankTransferStr
	case constants.CreditCard:
		return constants.CreditCardStr
	case constants.Cash:
		return constants.CashStr
	case constants.Gacha:
		return constants.GachaStr
	default:
		return ""
	}
}

func sourceOfFundsToInt(sof string) int {
	switch sof {
	case constants.WalletStr:
		return constants.Wallet
	case constants.BankTransferStr:
		return constants.BankTransfer
	case constants.CreditCardStr:
		return constants.CreditCard
	case constants.CashStr:
		return constants.Cash
	case constants.GachaStr:
		return constants.Gacha
	default:
		return 0
	}
}

func ToTransactionDTO(transaction entity.Transaction) TransactionDTO {
	return TransactionDTO{
		Id:                 transaction.Id,
		Type:               transaction.Type,
		SenderWalletNum:    transaction.SenderWalletNum,
		SenderName:         transaction.SenderName,
		RecepientWalletNum: transaction.RecepientWalletNum,
		RecepientName:      transaction.RecipientName,
		Amount:             transaction.Amount,
		SourceOfFunds:      sourceOfFundsToInt(transaction.SourceOfFunds),
		Description:        transaction.Description,
		Date:               transaction.Date.Format(constants.TimeStr),
	}
}

func ToTransactionsDTO(transactions []entity.Transaction) []TransactionDTO {
	transactionsDTO := []TransactionDTO{}

	for _, transaction := range transactions {
		transactionDTO := ToTransactionDTO(transaction)

		transactionsDTO = append(transactionsDTO, transactionDTO)
	}

	return transactionsDTO
}

func ToTransactionListDTO(transactionList entity.TransactionList) TransactionListDTO {
	return TransactionListDTO{
		TotalItem:    transactionList.TotalItem,
		Page:         transactionList.Page,
		Transactions: ToTransactionsDTO(transactionList.Transactions),
	}
}
