package main

import (
	"example.com/m/src/domain"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//用户名：密码@tcp(地址:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local
	var username string
	var password string
	var host string
	var dbname string
	fmt.Println("请依次输入mysql用户名 密码 地址:端口 数据库名 用空格分隔")
	fmt.Println("例如：root 123456 127.0.0.1:3306 test")
	fmt.Scanln(&username, &password, &host, &dbname)
	var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败，请检查您的输入及mysql配置")
	}

	// 迁移 schema
	db.AutoMigrate(&domain.AliveOrder{})
	db.AutoMigrate(&domain.AllStock{})
	db.AutoMigrate(&domain.StockInformation{})
	db.AutoMigrate(&domain.UserHolding{})
	db.AutoMigrate(&domain.UserInformation{})
	fmt.Println("数据库表迁移成功，输入任意键结束")
	fmt.Scanln()
}
