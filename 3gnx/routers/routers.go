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
		V1Group.POST("/manager/login", controller.MangerLogin)
		//4.注册时发送验证码
		V1Group.POST("/user/register-email", controller.SendEmailRegister)
		//5.修改密码验证邮箱时发送验证码
		V1Group.POST("/user/reset-email", controller.SendEmailReSet)
		//6.验证身份
		V1Group.POST("/user/VerifyCode-email", controller.ResetCodeVerify)
		//7.重设密码
		V1Group.POST("/user/reset-password", controller.ResetPassword)
	}
	return r
}
