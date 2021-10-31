package main

import(
	"example.com/m/src/dao"
	"example.com/m/src/domain"
)
func main()  {
	//user:=domain.UserInformation{
	//	UserID: "test",
	//	UserName: "test_name",
	//	FreeMoneyAmount: 5000.05,
	//	TotalMoneyAmount: 1000.05,
	//}
	var user domain.UserInformation
	dao.SearchOne(user,"user_id=?","test")
}
