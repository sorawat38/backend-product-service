package models

import (
	"github.com/shopspring/decimal"
)

type Menu struct {
	MenuID      int
	Name        string
	Description string
	DisplayPic  string
	Price       decimal.Decimal
}
