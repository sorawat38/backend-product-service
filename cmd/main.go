package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/config"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/adaptor/repositories/database/menudb"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/handler"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/handler/menuhdl"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/helper/logger"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/internal/service/menusrv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	logger.InitLog(cfg.Log)
	defer logger.CloseLogger()

	db, err := initDB(cfg.Database)
	if err != nil {
		panic(err)
	}

	// Init repository
	productDB := menudb.New(db)
	menuService := menusrv.New(productDB)
	menuHandler := menuhdl.NewHTTPHandler(&menuService)

	// Starting server
	e := echo.New()
	handler.InitRoute(e, menuHandler)

	log.Printf("Starting server on port %v...\n", cfg.App.Port)
	if err := e.Start(":" + cfg.App.Port); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func initDB(dbCfg config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DBName,
	)

	log.Printf("Connecting database... %v\n", dsn)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Connect database successfully!!")

	return gormDB, err
}
