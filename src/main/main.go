package main

import (
	"example.com/m/src/controller"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	//bind get function here
	router.GET("/ping", controller.Ping)
	//router.GET("/get/searchStock:stockName",controller.SearchStock)


	//bind post function here
	router.POST("/post/register", controller.Register)
	router.POST("/post/login", controller.Login)
	router.POST("/post/DrawK", controller.DrawK)
	router.POST("/post/buy", controller.BuyStock)
	router.POST("/post/sell",controller.SellStock)
	router.POST("/post/holdings", controller.SearchHoldings)


	router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}