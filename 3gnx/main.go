package main

import (
	"3gnx/dao"
	"3gnx/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 使用cors中间件，允许所有来源访问
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},              // 允许所有来源
		AllowMethods:     []string{"*"},              // 允许的请求方法
		AllowHeaders:     []string{"*"},              // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"}, // 允许暴露的响应头
		AllowCredentials: true,                       // 允许携带凭证（如Cookie）
	}))

	//创建连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() // 程序退出关闭数据库连接
	// 注册路由
	r := routers.SetUpRouter()
	r.Run(":8080")
}
