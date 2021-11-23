package main

import (
	"example.com/m/src/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

/*
func CrosHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "false")
		context.Set("content-type", "application/json")

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, result.Result{Code: result.OK, Data: "Options Request!"})
		}

		//处理请求
		context.Next()
	}
}*/

func main() {
	router := gin.Default()
	router.Use(cors.Default())
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