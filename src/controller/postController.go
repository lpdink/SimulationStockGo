package controller

import (
	"example.com/m/src/service"
	"github.com/gin-gonic/gin"
)
/*
post-controller要做的事情很简单，从前端取出数据，判断一下数据是否存在：
	如果某个数据不存在，可以使用默认值，如果是必须的数据，就返回失败。
	失败统一返回400，message信息可以自定。
取出数据，判断是否存在，将取出的数据传递给service层的业务函数。
根据service层业务函数的返回值，决定向前端的返回信息，无论业务处理成功与否，都返回200，在返回的json-message信息里填写处理结果。
具体写法参考func Register()
*/



/* 获取post请求参数的方法
//获取 id 参数, 通过 PostForm 获取的参数值是 String 类型。
id := c.PostForm("id")

// 跟 PostForm 的区别是可以通过第二个参数设置参数默认值
id := c.DefaultPostForm("id", "123")

// 获取 id 参数, 通过 GetPostForm 获取的参数值也是 String 类型,
// 区别是 GetPostForm 返回两个参数，第一个是参数值，第二个参数是参数是否存在的bool值，
// 可以用来判断参数是否存在。
id, ok := c.GetPostForm("id")
if !ok {
// 参数不存在
}
*/

func Register(c *gin.Context){

	//getData
	id, idOk :=c.GetPostForm("id")
	username, usernameOk :=c.GetPostForm("username")
	password, passwordOk :=c.GetPostForm("password")
	//checkIsExist
	if !idOk || !usernameOk ||!passwordOk {
		c.JSON(400,gin.H{
			"message":"value not found",
		})
		return
	}
	// call service function do something
	message:=service.Register(id,username,password)

	//return success and info to clienter
	c.JSON(200,gin.H{
		"message":message,
		"id":id,
		"username":username,
	})

}

//TODO 参考func Register，完成以下三个函数
//从前端提取的参数:
//参考https://github.com/lpdink/SimulationStockBasedOnQQBot/blob/master/controller/awesome/plugins/QQBotPlugins.py
func AddSelfStock(c *gin.Context)  {
	
}

func BuyStock(c *gin.Context)  {
	
}

func SellStock(c *gin.Context)  {

}

