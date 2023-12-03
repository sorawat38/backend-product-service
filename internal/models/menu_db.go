package models

import (
	"github.com/shopspring/decimal"
)

type Menu struct {
	MenuID      string
	Name        string
	Description string
	DisplayPic  string
	Price       decimal.Decimal
}
