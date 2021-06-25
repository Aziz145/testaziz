package infra

import (
	"github.com/Aziz145/testaziz/model"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func LoadMySQLDB() *gorm.DB {
	dsn := "root@tcp(127.0.0.1:3306)/user_manag?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&model.Users{}) 
	return db
}