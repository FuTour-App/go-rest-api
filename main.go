package main

import (
	"github.com/FuTour-App/go-rest-api/config"

	"github.com/FuTour-App/go-rest-api/models"
	"github.com/FuTour-App/go-rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func initConfig() *config.Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {

	}

	return config.NewConfig()
}

func main() {
	r := gin.Default()
	cfg := initConfig()
	models.ConnectDatabase(cfg)

	r.Static("/uploads", "./uploads")
	r.RedirectTrailingSlash = true
	routes.AuthRoutes(r, cfg)
	routes.ProductRoutes(r, cfg)

	r.Run()
}
