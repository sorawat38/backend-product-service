package database

import "github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"

type MenuRepository interface {
	GetAll() ([]models.Menu, error)
}
