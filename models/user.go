package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	FullName string `gorm:"varchar(255)" json:"full_name"`
	UserName string `gorm:"varchar(255)" json:"username"`
	Email    string `gorm:"varchar(255)" json:"email"`
	Password string `gorm:"varchar(255)" json:"password"`
}
