package models

import (
	"github.com/shopspring/decimal"
)

type Menu struct {
	MenuID int
	Name   string
	Price  decimal.Decimal
}
