package main

import (
	"example.com/m/src/controller"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	//bind get function here
	router.GET("/ping", controller.Ping)
	router.GET("/get/searchStock:stockName",controller.SearchStock)


	//bind post function here
	router.POST("/post/register", controller.Register)


	//bind delete function here


	//bind put function here


	router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}