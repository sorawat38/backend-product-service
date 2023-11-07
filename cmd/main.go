package main

import (
	"fmt"
	"log"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-product-service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	_, _ = initDB(cfg.Database)
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
