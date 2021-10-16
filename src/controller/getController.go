package controller


//get 请求的参数提取，可以参考：https://blog.csdn.net/qq_41004932/article/details/119486982

import (
	"example.com/m/src/service"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"ping success!",
	})
}

func SearchStock(c *gin.Context){
	stockName:=c.Param("stockName")
	if stockName==""{
		c.JSON(400,gin.H{
			"message":"value not found",
		})
		return
	}
	nowPrice,openPrice,closePrice :=service.SearchStock(stockName)
	c.JSON(200,gin.H{
		"nowPrice":nowPrice,
		"openPrice":openPrice,
		"closePrice":closePrice,
	})
}

//TODO 参考func SearchStock,完成以下三个函数
//从前端提取的参数:
//参考https://github.com/lpdink/SimulationStockBasedOnQQBot/blob/master/controller/awesome/plugins/QQBotPlugins.py
func SearchUserHoldings(c *gin.Context)  {

}

func SearchUserInformations(c *gin.Context){

}

//查询K线图
func SearchK(c *gin.Context)  {

}
