package entity

import "github.com/shopspring/decimal"

type Wallet struct {
	Id         int
	Number     string
	UserId     int
	Balance    decimal.Decimal
	Income     decimal.Decimal
	Expense    decimal.Decimal
	GachaTrial int
}
