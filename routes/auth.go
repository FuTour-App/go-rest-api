package routes

import (
	"github.com/FuTour-App/go-rest-api/config"
	"github.com/FuTour-App/go-rest-api/controllers/authcontroller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, cfg *config.Config) {
	r.POST("/register", authcontroller.Register)
	r.POST("/login", func(c *gin.Context) {
		authcontroller.Login(c, cfg)
	})
	r.POST("/logout", authcontroller.Logout)
}
