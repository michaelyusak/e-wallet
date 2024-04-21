package entity

import "github.com/shopspring/decimal"

type GachaBox struct {
	BoxNumber int
	PrizeId   int
	Prize     decimal.Decimal
}
