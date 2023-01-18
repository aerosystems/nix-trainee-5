package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewClient() *gorm.DB {
	dsn := "sandbox_user:passpass@tcp(127.0.0.1:3306)/sandbox?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
