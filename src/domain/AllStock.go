package domain

type AllStock struct {
	StockID   string `gorm:"column:stock_id;primary_key"`
	StockName string `gorm:"column:stock_name;NOT NULL"`
}