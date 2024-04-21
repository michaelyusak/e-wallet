package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Id                 int
	Type               string
	SenderWalletNum    string
	SenderName         string
	RecepientWalletNum string
	RecipientName      string
	Amount             decimal.Decimal
	SourceOfFunds      string
	Description        string
	Date               time.Time
}

type TransactionList struct {
	TotalItem    int
	Page         string
	Transactions []Transaction
}
