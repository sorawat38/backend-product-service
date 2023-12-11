package menudb

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/adaptor/repositories/database"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) database.MenuRepository {
	return menuRepo{
		db: db,
	}
}

func (repo menuRepo) GetById(id string) (models.Menu, error) {

	var menus models.Menu
	tx := repo.db.Where(models.Menu{MenuID: id}).Find(&menus)
	if tx.Error != nil {
		return models.Menu{}, tx.Error
	}

	return menus, nil
}

func (repo menuRepo) GetAll() ([]models.Menu, error) {

	var menus []models.Menu
	tx := repo.db.Find(&menus)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return menus, nil
}
