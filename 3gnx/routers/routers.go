package routers

import (
	"3gnx/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))

	// 注册 session 中间件到路由组

	//store := cookie.NewStore([]byte("your-secret-key"))
	//r.Use(sessions.Sessions("session", store))
	V1Group := r.Group("api")
	{
		V1Group.Use(sessions.Sessions("session", store))
		//middles.GetSessionId(V1Group)
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
		//8.学生报名
		V1Group.POST("/user/login/apply", controller.StudentApplySuccess)
		//9.查询一面通过学生
		V1Group.GET("/manager/login/getOne", controller.GetApplyStuListOne)
		//10.查询二面通过学生
		V1Group.GET("/manager/login/getTwo", controller.GetApplyStuListTwo)
		//11.学生查询自己一面状况
		V1Group.POST("/user/login/getOne", controller.StuGetApplyStuListOne)
		//12.学生查询自己er面状况
		V1Group.POST("/user/login/getTwo", controller.StuGetApplyStuListTwo)
		//13.设置一面通过
		V1Group.PUT("/manager/login/SetOne", controller.SetOneSuccessfully)
		//14.设置二面通过
		V1Group.PUT("/manager/login/SetTwo", controller.SetTwoSuccessfully)
	}
	return r
}
