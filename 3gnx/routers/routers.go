package routers

import (
	"3gnx/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	V1Group := r.Group("api")
	{
		//待办事项
		//1.用户注册(包含添加用户)
		V1Group.POST("/user/register", controller.UserRegister)
		//2.用户登录
		V1Group.POST("/user/login", controller.UserLogin)
		//3.管理员登录
		//4.发送验证码
	}
	return r
}
