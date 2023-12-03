package models

type MenuGetAllResponse struct {
	CommonResponse
	Data []MenuGetAllResponseBody `json:"data,omitempty"`
}

type MenuGetAllResponseBody struct {
	Id          string  `json:"id"`
	FNname      string  `json:"fName"`
	Description string  `json:"desc,omitempty"`
	DisplayPic  string  `json:"displayPic"`
	Price       float64 `json:"price"`
}

type MenuGetByIdResponse struct {
	CommonResponse
	Data MenuGetByIdResponseBody `json:"data,omitempty"`
}

type MenuGetByIdResponseBody struct {
	Id          string  `json:"id"`
	FNname      string  `json:"fName"`
	Description string  `json:"desc,omitempty"`
	DisplayPic  string  `json:"displayPic"`
	Price       float64 `json:"price"`
}
