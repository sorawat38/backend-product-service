package menusrv

import (
	"errors"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/adaptor/repositories/database"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"
)

type service struct {
	menuRepo database.MenuRepository
}

func New(menuRepo database.MenuRepository) service {
	return service{
		menuRepo: menuRepo,
	}
}

func (srv *service) GetById(id int) (models.Menu, error) {

	result, err := srv.menuRepo.GetById(id)
	if err != nil {
		return models.Menu{}, err
	}

	return result, nil
}

func (srv *service) GetAll() ([]models.Menu, error) {

	resultList, err := srv.menuRepo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(resultList) == 0 {
		return nil, errors.New("can't find any flavours")
	}

	return resultList, nil
}
