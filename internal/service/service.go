package service

import "github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"

type MenuService interface {
	GetById(id int) (models.Menu, error)
	GetAll() ([]models.Menu, error)
}
