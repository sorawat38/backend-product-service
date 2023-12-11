package menuhdl

import (
	"net/http"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/constant"
	apiresponses "github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/handler/apiResponses"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/service"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	menuService service.MenuService
}

func NewHTTPHandler(menuService service.MenuService) HTTPHandler {
	return HTTPHandler{
		menuService: menuService,
	}
}

func (hdl *HTTPHandler) GetById(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, apiresponses.InvalidInputError(nil))
	}

	// newId, err := strconv.Atoi(id)
	// if err != nil {

	// 	return echo.NewHTTPError(http.StatusBadRequest, apiresponses.InvalidInputError(err))
	// }

	result, err := hdl.menuService.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, apiresponses.InternalError(err))
	}

	var response models.MenuGetByIdResponse
	response.Code = constant.SuccessCode
	response.Message = constant.SuccessMessage

	var respData models.MenuGetByIdResponseBody
	respData.Id = result.MenuID
	respData.FNname = result.Name
	respData.Description = result.Description
	respData.DisplayPic = result.DisplayPic
	price, _ := result.Price.Float64()
	respData.Price = price
	response.Data = respData

	return c.JSON(http.StatusOK, response)
}

func (hdl *HTTPHandler) GetAll(c echo.Context) error {

	var response models.MenuGetAllResponse
	resultList, err := hdl.menuService.GetFlavorsAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	var responseBodyList = make([]models.MenuGetAllResponseBody, 0, len(resultList))
	for _, each := range resultList {
		var responseBody models.MenuGetAllResponseBody
		responseBody.Id = each.MenuID
		responseBody.FNname = each.Name
		responseBody.Description = each.Description
		responseBody.DisplayPic = each.DisplayPic
		price, _ := each.Price.Float64()
		responseBody.Price = price
		responseBodyList = append(responseBodyList, responseBody)
	}

	response.Code = constant.SuccessCode
	response.Message = constant.SuccessMessage
	response.Data = responseBodyList

	return c.JSON(http.StatusOK, response)
}
