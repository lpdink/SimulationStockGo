package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"


func SearchOne(table interface{}, key string, value string){
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println(key, value)
	db.First(&table, key, value)
	log.Println(table)
}
