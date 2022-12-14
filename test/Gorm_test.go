package test

import (
	"ginchat/models"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGorm(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	// user := &models.UserBasic{}
	// user.Name = "申专"
	// db.Create(user)
	// // Read
	// fmt.Println(db.First(user, 1)) // 根据整型主键查找

	// // Update - 将 product 的 price 更新为 200
	// db.Model(user).Update("PassWord", 1234)

}
