package main

import (
	"gogochat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	connstr := "root:123456@tcp(127.0.0.1:3306)/gogochat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connstr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.UserBasic{})
}
