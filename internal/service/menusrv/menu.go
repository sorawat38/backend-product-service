package menusrv

import (
	"errors"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/adaptor/repositories/database"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/helper/logger"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/models"
	"go.uber.org/zap"
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
		logger.Error("can't find menu by id", zap.Int("menu_id", id), zap.Error(err))
		return models.Menu{}, err
	}

	return result, nil
}

func (srv *service) GetAll() ([]models.Menu, error) {

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
