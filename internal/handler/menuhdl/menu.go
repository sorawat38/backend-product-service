package menuhdl

import (
	"net/http"
	"strconv"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/constant"
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
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	newId, err := strconv.Atoi(id)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	result, err := hdl.menuService.GetById(newId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	var response models.MenuGetByIdResponse
	response.Code = constant.SuccessCode
	response.Message = constant.SuccessMessage
	response.Data = models.MenuGetByIdResponseBody(result)

	return c.JSON(http.StatusOK, response)
}

func (hdl *HTTPHandler) GetAll(c echo.Context) error {

	var response models.MenuGetAllResponse
	resultList, err := hdl.menuService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	var responseBodyList = make([]models.MenuGetAllResponseBody, 0, len(resultList))
	for _, each := range resultList {
		responseBodyList = append(responseBodyList, models.MenuGetAllResponseBody(each))
	}

	response.Code = constant.SuccessCode
	response.Message = constant.SuccessMessage
	response.Data = responseBodyList

	return c.JSON(http.StatusOK, response)
}
