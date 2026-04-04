package authcontroller

import (
	"net/http"
	"time"

	"github.com/FuTour-App/go-rest-api/config"
	"github.com/FuTour-App/go-rest-api/helper"
	"github.com/FuTour-App/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var existingEmail models.User

	err := models.DB.Where("email = ?", user.Email).First(&existingEmail).Error

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "Email sudah digunakan"})
		return
	}

	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Gagal memproses password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Gagal menambahkan user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Registrasi berhasil", "user": user})

}

func Login(c *gin.Context, cfg *config.Config) {

	var user models.User
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "email atau password salah"})

		default:
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		}

	}

	match := helper.CheckPasswordHash(input.Password, user.Password)

	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "email atau password salah"})
		return
	}

	expTime := time.Now().Add(time.Minute * 20)
	claims := &config.JWTClaim{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.App.JwtIssuer,
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString([]byte(cfg.App.JwtSecretKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat token"})
		return
	}

	c.SetCookie(
		"token",
		token,
		3600*24,
		"/",
		"localhost",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Login Berhasil!"})

}

func Logout(c *gin.Context) {
	c.SetCookie(
		"token",
		"",
		3600*24,
		"/",
		"localhost",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Logout Berhasil!"})

}
