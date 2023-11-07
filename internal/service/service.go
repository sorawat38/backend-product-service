package service

import "github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"

type MenuService interface {
	GetAll() ([]models.Menu, error)
}
