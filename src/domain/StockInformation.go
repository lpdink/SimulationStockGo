package domain

import "time"

type StockInformation struct {
	StockID    string    `gorm:"column:stock_id;primary_key"`
	StockName  string    `gorm:"column:stock_name"`
	NowPrice   string    `gorm:"column:now_price"`
	FlushTime  time.Time `gorm:"column:flush_time"`
	UpDownRate float64   `gorm:"column:up_down_rate"`
}
