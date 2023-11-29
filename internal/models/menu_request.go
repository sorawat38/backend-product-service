package models

import "github.com/shopspring/decimal"

type MenuGetAllResponse struct {
	CommonResponse
	Data []MenuGetAllResponseBody `json:"data,omitempty"`
}

type MenuGetAllResponseBody struct {
	MenuID int             `json:"menu_id,omitempty"`
	Name   string          `json:"name,omitempty"`
	Price  decimal.Decimal `json:"price,omitempty"`
}

type MenuGetByIdResponse struct {
	CommonResponse
	Data MenuGetByIdResponseBody `json:"data,omitempty"`
}

type MenuGetByIdResponseBody struct {
	MenuID int             `json:"menu_id,omitempty"`
	Name   string          `json:"name,omitempty"`
	Price  decimal.Decimal `json:"price,omitempty"`
}
