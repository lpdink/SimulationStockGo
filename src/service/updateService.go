package service

import (
	"encoding/json"
	"example.com/m/src/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func getPrice(stockid string)(float64){
	params := map[string]interface{}{
		"ts_code":stockid,
	}
	var result = callAPI("daily","",params)
	var msg TushareMsg
	_ = json.Unmarshal(result, &msg)
	log.Println(msg)
	return msg.Data.Items[len(msg.Data.Items)-1][5].(float64)
}

func BuyStock(stockid string, stock_name string, userid string, stocknum int)(int){
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		return -1
	}
	price:=getPrice(stockid)
	total_price:=price*float64(stocknum)
	var user = domain.UserInformation{}
	db.Where("user_id=?",userid).First(&user)
	if total_price>user.FreeMoneyAmount{
		return -1
	} else {
		var holding =domain.UserHolding{}
		tx:=db.Where("stock_id=?",stockid)
		tx=tx.Where("user_id=?",userid)
		var is_exist =tx.First(&holding).Error==nil
		if is_exist{
			holding.StockAmount+=int(stocknum)
			db.Model(&holding).Update("stock_amount", holding.StockAmount)
			log.Println("enter is exist!!!!!")
		} else{
			db.Create(domain.UserHolding{
				StockID:          stockid,
				UserID:           userid,
				StockName:        stock_name,
				StockAmount:      int(stocknum),
				BoughtPrice:      "",
				BoughtTotalPrice: 0,
				BoughtTime:       time.Now(),
			})
		}
		user.FreeMoneyAmount-=total_price
		db.Model(&user).Update("free_money_amount",user.FreeMoneyAmount)
	}
	return 1

}

func SellStock(stockid string, userid string, stocknum int) int {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		return -1
	}
	price:=getPrice(stockid)
	total_price:=price*float64(stocknum)
	var holding =domain.UserHolding{}
	tx:=db.Where("id=?",stockid)
	tx=tx.Where("user_id=?",userid)
	var user = domain.UserInformation{}
	db.Where("user_id=?",userid).First(&user)
	var not_exist =tx.First(&holding).Error!=nil
	if not_exist{
		return -1
	}else {
		if stocknum>holding.StockAmount{
			return -1
		}
		holding.StockAmount-=stocknum
		if holding.StockAmount==0{
			db.Delete(&holding)
			db.Model(&user).Update("free_money_amount", user.FreeMoneyAmount+total_price)
		}else {
			db.Model(&holding).Update("stock_amount", holding.StockAmount)
			db.Model(&user).Update("free_money_amount", user.FreeMoneyAmount+total_price)
		}
	}
	return 1
}
