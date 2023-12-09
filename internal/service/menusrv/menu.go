package menusrv

import (
	"errors"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/adaptor/repositories/database"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/helper/logger"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/service"
	"go.uber.org/zap"
)

type menuSrv struct {
	menuRepo database.MenuRepository
}

func New(menuRepo database.MenuRepository) service.MenuService {
	return menuSrv{
		menuRepo: menuRepo,
	}
}

func (srv menuSrv) GetById(id string) (models.Menu, error) {

	result, err := srv.menuRepo.GetById(id)
	if err != nil {
		logger.Error("can't find menu by id", zap.String("menu_id", id), zap.Error(err))
		return models.Menu{}, err
	}

	return result, nil
}

func (srv menuSrv) GetAll() ([]models.Menu, error) {

	resultList, err := srv.menuRepo.GetAll()
	if err != nil {
		logger.Error("can't find all menu", zap.Error(err))
		return nil, err
	}

	if len(resultList) == 0 {
		logger.Error("menu is empty")
		return nil, errors.New("can't find any flavours")
	}

	return resultList, nil
}
