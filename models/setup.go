package models

import (
	"fmt"

	"github.com/FuTour-App/go-rest-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *config.Config) {
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.User,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi ke database: " + err.Error())
	}

	database.AutoMigrate(&Product{})
	database.AutoMigrate(&User{})

	DB = database
}
