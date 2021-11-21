package domain

import "time"

type UserHolding struct {
	ID               string    `gorm:"column:id;primary_key"`
	UserID           string    `gorm:"column:user_id"`
	StockName        string    `gorm:"column:stock_name;NOT NULL"`
	StockAmount      int       `gorm:"column:stock_amount"`
	BoughtPrice      string    `gorm:"column:bought_price"`
	BoughtTotalPrice float64   `gorm:"column:bought_total_price"`
	BoughtTime       time.Time `gorm:"column:bought_time"`
}

