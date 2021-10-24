package main

import "fmt"

func main() {
	var username string
	var password string
	var host string
	var dbname string
	fmt.Println("请依次输入mysql用户名 密码 地址:端口 数据库名 用空格分隔")
	fmt.Println("例如：root 123456 127.0.0.1:3304 test")
	fmt.Scanln(&username, &password, &host, &dbname)
	var dns = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, dbname)
	fmt.Println(dns)
}
