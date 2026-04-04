package main

import (
	"github.com/FuTour-App/go-rest-api/config"
	"github.com/FuTour-App/go-rest-api/controllers/authcontroller"
	"github.com/FuTour-App/go-rest-api/controllers/productcontroller"
	"github.com/FuTour-App/go-rest-api/middleware"
	"github.com/FuTour-App/go-rest-api/models"
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

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)

	authorized := r.Group("/api/products")

	authorized.Use(middleware.AuthMiddleWare(cfg))
	{
		authorized.POST("/", productcontroller.Create)
		authorized.PUT("/:id", productcontroller.Update)
		authorized.DELETE("/:id", productcontroller.Delete)
	}

	r.POST("/register", authcontroller.Register)
	r.POST("/login", func(c *gin.Context) {
		authcontroller.Login(c, cfg)
	})
	r.POST("/logout", authcontroller.Logout)

	r.Run()
}
