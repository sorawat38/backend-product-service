package database

import "github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"

type MenuRepository interface {
	GetById(id int) (models.Menu, error)
	GetAll() ([]models.Menu, error)
}
