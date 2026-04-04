package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"type:varchar(300)" json:"nama_product"`
	Deskripsi   string `gorm:"type:text" json:"deskripsi"`
	Image       string `gorm:"type:varchar(255)" json:"image"`
	Stock       int    `gorm:"type:int" json:"stock"`
	Price       int64  `gorm:"type:bigint" json:"price"`
}
