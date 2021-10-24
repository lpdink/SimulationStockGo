package domain

type UserInformation struct {
	UserID           string  `gorm:"column:user_id;primary_key"`
	UserName         string  `gorm:"column:user_name"`
	FreeMoneyAmount  float64 `gorm:"column:free_money_amount"`
	TotalMoneyAmount float64 `gorm:"column:total_money_amount"`
}
