package dto

import "github.com/shopspring/decimal"

type WalletDTO struct {
	Id         int             `json:"id"`
	Number     string          `json:"number"`
	Balance    decimal.Decimal `json:"balance"`
	Income     decimal.Decimal `json:"income"`
	Expense    decimal.Decimal `json:"expense"`
	GachaTrial int             `json:"gacha_trial"`
}
