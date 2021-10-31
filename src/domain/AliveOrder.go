package domain

import "time"

type AliveOrders struct {
	UserID           string    `gorm:"column:user_id;primary_key"`
	AliveOrderIndex  int       `gorm:"column:alive_order_index;NOT NULL"`
	AliveOrderTime   time.Time `gorm:"column:alive_order_time"`
	BuyOrSell        int       `gorm:"column:buy_or_sell"`
	StockID          string    `gorm:"column:stock_id"`
	StockName        string    `gorm:"column:stock_name"`
	StockPrice       string    `gorm:"column:stock_price"`
	OrderMoneyAmount float64   `gorm:"column:order_money_amount"`
	StockAmount      int       `gorm:"column:stock_amount"`
	IsAlive          int       `gorm:"column:is_alive"`
}