package menuhdl

import (
	"net/http"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type Response struct {
	models.CommonResponse
	Data []ResponseBody `json:"data,omitempty"`
}

type ResponseBody struct {
	MenuID     int             `json:"menu_id,omitempty"`
	Name       string          `json:"name,omitempty"`
	Ingredient string          `json:"ingredient,omitempty"`
	Price      decimal.Decimal `json:"price,omitempty"`
}

type HTTPHandler struct {
	menuService service.MenuService
}

func NewHTTPHandler(menuService service.MenuService) HTTPHandler {
	return HTTPHandler{
		menuService: menuService,
	}
}

func (hdl *HTTPHandler) GetAll(c echo.Context) error {

	var response Response
	resultList, err := hdl.menuService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	var responseBodyList = make([]ResponseBody, 0, len(resultList))
	for _, each := range resultList {
		responseBodyList = append(responseBodyList, ResponseBody(each))
	}

	response.Code = 100
	response.Message = "Successfully"
	response.Data = responseBodyList

	return c.JSON(http.StatusOK, response)
}
