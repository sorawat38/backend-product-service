package menudb

import (
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) repository {
	return repository{
		db: db,
	}
}

func (repo repository) GetById(id string) (models.Menu, error) {

	var menus models.Menu
	tx := repo.db.Where(models.Menu{MenuID: id}).Find(&menus)
	if tx.Error != nil {
		return models.Menu{}, tx.Error
	}

	return menus, nil
}

func (repo repository) GetAll() ([]models.Menu, error) {

	var menus []models.Menu
	tx := repo.db.Find(&menus)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return menus, nil
}
