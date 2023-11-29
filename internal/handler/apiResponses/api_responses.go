package apiresponses

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/constant"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"
)

func SuccessResponse() models.CommonResponse {
	return models.CommonResponse{
		Code:    constant.SuccessCode,
		Message: constant.SuccessMessage,
	}
}

func InvalidInputError() models.CommonResponse {
	return models.CommonResponse{
		Code:    constant.InvalidInputCode,
		Message: constant.InvalidInputMsg,
	}
}

func InternalError() models.CommonResponse {
	return models.CommonResponse{
		Code:    constant.InternalErrorCode,
		Message: constant.InternalError,
	}
}
