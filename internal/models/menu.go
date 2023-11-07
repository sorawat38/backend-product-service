package models

import (
	"github.com/shopspring/decimal"
)

type Menu struct {
	MenuID     int
	Name       string
	Ingredient string
	Price      decimal.Decimal
}
