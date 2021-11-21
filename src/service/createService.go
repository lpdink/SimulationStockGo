package service

import (
	"example.com/m/src/domain"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var username = "root"
var password = "123456"
var host = "127.0.0.1:3306"
var dbname = "test"
var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, dbname)

func Register(userid string, username string, password string) (string,int) {
	var user = domain.UserInformation{}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return "连接数据库失败",0
	}
	var is_exist = db.Where("user_id=?",userid).First(&user).Error==nil
	if is_exist{
		return "您已经注册",0
	} else {
		var new_user = domain.UserInformation{
			UserID:           userid,
			UserName:         username,
			PassWord: password,
			FreeMoneyAmount:  500000,
			TotalMoneyAmount: 500000,
		}
		db.Create(new_user)
		return "ok",1
	}
}


