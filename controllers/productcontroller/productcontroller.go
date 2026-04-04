package productcontroller

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/FuTour-App/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var products []models.Product

	name := c.Query("nama_product")
	query := models.DB

	if name != "" {
		query = query.Where("nama_product LIKE ?", "%"+name+"%")
	}

	if err := query.Find(&products).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})

}

func Show(c *gin.Context) {

	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": product})

}
func Create(c *gin.Context) {

	if err := c.Request.ParseMultipartForm(8 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "File harus lebih kecil dari 2MB, dan berformat jpg/jpeg/png"})
		return
	}

	namaProduct := c.PostForm("nama_product")
	deskripsi := c.PostForm("deskripsi")
	stock := c.PostForm("stock")
	price := c.PostForm("price")

	if namaProduct == "" || price == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Nama produk dan harga wajib diisi"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Gambar wajib diupload"})
		return
	}

	extension := strings.ToLower(filepath.Ext(file.Filename))
	if extension != ".jpg" && extension != ".jpeg" && extension != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Format file harus jpg/jpeg/png"})
		return
	}

	if file.Size > 2*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Ukuran file maksimal 2MB"})
		return
	}

	filename := time.Now().Format("20060102150405") + "-" + file.Filename
	path := "uploads/" + filename

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan gambar ke server"})
		return
	}

	priceInt, _ := strconv.ParseInt(price, 10, 64)
	stockInt, _ := strconv.Atoi(stock)

	product := models.Product{
		NamaProduct: namaProduct,
		Deskripsi:   deskripsi,
		Image:       filename,
		Stock:       stockInt,
		Price:       priceInt,
	}

	if err := models.DB.Create(&product).Error; err != nil {

		os.Remove(path)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan data ke database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Produk baru berhasil ditambahkan",
		"product": product,
	})
}

func Update(c *gin.Context) {

	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Produk tidak ditemukan"})
		return
	}

	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request tidak valid atau file terlalu besar"})
		return
	}

	namaProduct := c.PostForm("nama_product")
	deskripsi := c.PostForm("deskripsi")
	stock := c.PostForm("stock")
	price := c.PostForm("price")

	file, err := c.FormFile("image")

	if err == nil {

		extension := strings.ToLower(filepath.Ext(file.Filename))
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Format file baru harus jpg/jpeg/png"})
			return
		}

		if file.Size > 2*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Ukuran file baru maksimal 2MB"})
			return
		}

		newFilename := time.Now().Format("20060102150405") + "-" + file.Filename
		newPath := "uploads/" + newFilename

		if err := c.SaveUploadedFile(file, newPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan gambar baru"})
			return
		}

		if product.Image != "" {
			os.Remove("uploads/" + product.Image)
		}

		product.Image = newFilename
	}

	priceInt, _ := strconv.ParseInt(price, 10, 64)
	stockInt, _ := strconv.Atoi(stock)

	product.NamaProduct = namaProduct
	product.Deskripsi = deskripsi
	product.Price = priceInt
	product.Stock = stockInt

	if err := models.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui data di database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Produk berhasil diperbarui",
		"product": product,
	})

}

func Delete(c *gin.Context) {

	var product models.Product
	id := c.Param("id")

	if models.DB.Model(&product).Where("id = ?", id).Delete(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat menghapus produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "produk berhasil dihapus"})

}
