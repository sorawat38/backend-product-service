package models

import "github.com/shopspring/decimal"

type MenuGetAllResponse struct {
	CommonResponse
	Data []MenuGetAllResponseBody `json:"data,omitempty"`
}

type MenuGetAllResponseBody struct {
	Id          int             `json:"id"`
	FNname      string          `json:"fNname"`
	Description string          `json:"desc,omitempty"`
	DisplayPic  string          `json:"displayPic"`
	Price       decimal.Decimal `json:"price"`
}

type MenuGetByIdResponse struct {
	CommonResponse
	Data MenuGetByIdResponseBody `json:"data,omitempty"`
}

type MenuGetByIdResponseBody struct {
	Id          int             `json:"id"`
	FNname      string          `json:"fNname"`
	Description string          `json:"desc,omitempty"`
	DisplayPic  string          `json:"displayPic"`
	Price       decimal.Decimal `json:"price"`
}
