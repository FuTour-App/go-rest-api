package routes

import (
	"github.com/FuTour-App/go-rest-api/config"
	"github.com/FuTour-App/go-rest-api/controllers/productcontroller"
	"github.com/FuTour-App/go-rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, cfg *config.Config) {
	r.GET("api/products", productcontroller.Index)
	r.GET("api/products/:id", productcontroller.Show)

	authorized := r.Group("/api/products")
	authorized.Use(middleware.AuthMiddleWare(cfg))
	{
		authorized.POST("", productcontroller.Create)
		authorized.PUT("/:id", productcontroller.Update)
		authorized.DELETE("/:id", productcontroller.Delete)
	}
}
