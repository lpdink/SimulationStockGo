package service

import "github.com/gin-gonic/gin"

func SearchStock(stockName string) (string, string, string) {
	//TODO call python function
	return stockName, "", ""
}

//TODO call dao function finish the job
func SearchUserHoldings(userId string) (string, []gin.H) {
	// TODO
	return userId, nil
}

func SearchUserInformations(userId string) (string, string, float64, float64) {
	// TODO
	return userId, "", 0.0, 0.0
}

//查询K线图
func SearchK() {

}
