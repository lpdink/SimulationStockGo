# SimulationStockGo
## 概述
<b>基于golong-Gin的模拟金融股票交易系统的HTTP-Server，使用redis和mysql.</b>
## 项目结构：MVC
- src: 源代码
    - controller: 与前端进行交互
      - getController: 实现get方法
      - postController: 实现post方法
      - deleteController: 实现delete方法
      - putController: 实现put方法
    - service: 业务处理层
      - createService: 实现创建相关业务
      - deleteService: 实现删除相关业务
      - updateService: 实现更新相关业务
      - searchService: 实现查询相关业务
    - dao: 数据库交互层
      - TODO
    - domain: 数据库实体类，用于orm方法
      - TODO
    - main: 绑定router处理函数，启动server
- resource: 资源文件
- test: 测试用例，api测试等
## 数据结构设计
- 参考resource文件下的simulationstock.sql文件
## TODO
- 完成getController.go中的三个函数
- 完成postController.go中的三个函数
- 参考[SimulationStockBasedOnQQBot](https://github.com/lpdink/SimulationStockBasedOnQQBot) 
  中service层的代码，实现本工程service层的业务逻辑。
    - Notes: tushare是python的，在go中需要手写HTTP请求去调用，参考https://tushare.pro/document/1?doc_id=130。    
- 查询gorm资料，在实现domain实体类后，实现dao层增删改查函数。
## DONE
- 完成了项目基本框架的建立
## 依赖说明
- [goLand: IDE环境](https://www.jetbrains.com/go/) go-sdk==1.16.2
- [gin: web框架](https://github.com/gin-gonic/gin)
- [gorm: orm框架](https://gorm.io/zh_CN/docs/index.html)
## 依赖配置
- 下载goLand社区版,配置好go环境
- 设置依赖下载的代理
```go
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
//执行完毕后，可以通过go env命令查看参数是否被正确修改了。
//如果没有修改，尝试重启cmd/Terminal/终端
```
- 初始化项目go mod(git 仓库中已经包含了go.mod文件)
```go
go mod init example.com/m
```
- 下载依赖
```go
//gin
go get -u github.com/gin-gonic/gin
//gorm
go get -u gorm.io/gorm 
go get -u gorm.io/driver/sqlite
```
- 初始化项目MySQL表
```shell
./tools/loadMySql.exe
# 根据命令行提示，输入信息即可。
```
