package middleware

import (
	"net/http"

	"github.com/FuTour-App/go-rest-api/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleWare(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Sesi berakhir, silahkan login"})
			c.Abort()
			return
		}
		claims := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.App.JwtSecretKey), nil
		})

		if err != nil || !token.Valid || claims.Issuer != cfg.App.JwtIssuer {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid atau kadaluarsa"})
			c.Abort()
		}
		c.Set("user_email", claims.Email)
		c.Next()

	}
}
