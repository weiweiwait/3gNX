package controller

import (
	"3gnx/models"
	"3gnx/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var newsession string

// 用户注册并返回信息给客户端
func UserRegister(c *gin.Context) {

	sessionID, _ := c.Cookie("session")
	//var requestData struct {
	//	Username string `json:"username"`
	//	Password string `json:"password"`
	//	Class    string `json:"class"`
	//	Xuehao   string `json:"xuehao"`
	//	Email    string `json:"email"`
	//	Code     string `json:"code"`
	//}
	var requestData struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
		Class    string `form:"class" json:"class"`
		Xuehao   string `form:"xuehao" json:"xuehao"`
		Email    string `form:"email" json:"email"`
		Code     string `form:"code" json:"code"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用UserRegister函数进行注册
	result := server.UserRegister(requestData.Username, requestData.Password, requestData.Class, requestData.Xuehao, requestData.Email, requestData.Code, sessionID)

	//// 根据注册结果返回相应的数据给前端
	//if result == "" {
	//	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
	//} else {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": result})
	//}
	// 封装注册结果为RestBean对象
	var restBeanRegister *models.RestBean
	if result == "" {
		restBeanRegister = models.SuccessRestBeanWithData("注册成功")

	} else {
		restBeanRegister = models.FailureRestBeanWithData(http.StatusBadRequest, result)
	}
	//返回注册结果给前端
	c.JSON(restBeanRegister.Status, restBeanRegister)
}

// 用户登录并返回信息给客户端
func UserLogin(c *gin.Context) {
	var requestData struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	var restBeanLogin *models.RestBean
	result := server.UserLogin(requestData.Username, requestData.Password)
	if result == "" {
		restBeanLogin = models.SuccessRestBeanWithData("登录成功")
	} else {
		restBeanLogin = models.FailureRestBeanWithData(http.StatusBadRequest, result)
	}
	//返回登录结果给前端
	c.JSON(restBeanLogin.Status, restBeanLogin)
}

// 管理员登录并返回信息给客户端
func MangerLogin(c *gin.Context) {
	var requestData struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	var restBeanLogin *models.RestBean
	result := server.ManagerLogin(requestData.Username, requestData.Password)
	if result == "" {
		restBeanLogin = models.SuccessRestBeanWithData("登录成功")
	} else {
		restBeanLogin = models.FailureRestBeanWithData(http.StatusBadRequest, result)
	}
	//返回登录结果给前端
	c.JSON(restBeanLogin.Status, restBeanLogin)
}

// 注册时发送验证码
func SendEmailRegister(c *gin.Context) {
	sessionID, _ := c.Cookie("session")
	fmt.Println(sessionID)
	//sessionID := middles.GetSessionId(c)
	var requestData struct {
		Email string `form:"email" json:"email"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	// 调用UserRegister函数进行注册
	result := server.SendEmail(requestData.Email, sessionID, false)

	// 根据注册结果返回相应的数据给前端
	// 封装注册结果为RestBean对象
	var restBeanRegister *models.RestBean
	if result == "" {
		restBeanRegister = models.SuccessRestBeanWithData("邮件已发送，请注意查收")

	} else {
		restBeanRegister = models.FailureRestBeanWithData(http.StatusBadRequest, result)
	}
	//返回注册结果给前端
	c.JSON(restBeanRegister.Status, restBeanRegister)

}

// 修改密码时发送验证码
func SendEmailReSet(c *gin.Context) {
	sessionID, _ := c.Cookie("session")
	//sessionID := middles.GetSessionId(c)
	var requestData struct {
		Email string `form:"email" json:"email"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	// 调用UserRegister函数进行注册
	result := server.SendEmail(requestData.Email, sessionID, true)

	// 根据注册结果返回相应的数据给前端
	// 封装注册结果为RestBean对象
	var restBeanRegister *models.RestBean
	if result == "" {
		restBeanRegister = models.SuccessRestBeanWithData("邮件已发送，请注意查收")

	} else {
		restBeanRegister = models.FailureRestBeanWithData(http.StatusBadRequest, result)
	}
	//返回注册结果给前端
	c.JSON(restBeanRegister.Status, restBeanRegister)

}

// 邮箱身份验证，然后才能修改密码
func ResetCodeVerify(c *gin.Context) {
	//sessionID := middles.GetSessionId(c)
	sessionID, _ := c.Cookie("session")
	// 获取验证参数
	var requestData struct {
		Email string `form:"email" json:"email"`
		Code  string `form:"code" json:"code"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	// 调用函数进行身份验证
	result := server.ResetCode(requestData.Email, requestData.Code, sessionID)

	// 根据注册结果返回相应的数据给前端
	// 封装注册结果为RestBean对象
	var restBeanRegister *models.RestBean
	if result == "" {
		// 在会话中设置重置密码的相关属性
		//middles.SetSessionAttribute(c, "reset-password", requestData.Email)
		restBeanRegister = models.SuccessRestBean()

	} else {
		restBeanRegister = models.FailureRestBeanWithData(http.StatusBadRequest, result)
	}
	//返回结果给前端
	c.JSON(restBeanRegister.Status, restBeanRegister)

}

// 修改密码
func ResetPassword(c *gin.Context) {

	//sessionID, _ := c.Cookie("session")
	// 获取重置密码参数
	var requestData struct {
		Password string `form:"password" json:"password"`
		Email    string `form:"email" json:"email"`
	}

	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 从会话中获取重置密码的属性
	//email := middles.GetSessionAttribute(c, "reset-password")
	// 将邮箱地址转换为字符串
	//emailString, _ := email.(string)
	//var emailinit  =  strings.Index(newsession,"reset-password")
	//for i:=emailinit+14;i<len(sessionID)
	var restBeanRegister *models.RestBean
	//if email == nil {
	//	restBeanRegister = models.FailureRestBeanWithData(http.StatusBadRequest, "清先验证邮箱身份")
	status, _ := models.GetStatusByEmail(requestData.Email)
	fmt.Println(status)
	if status == 1 {
		if server.ResetPassword(requestData.Password, requestData.Email) == "" {
			//middles.DeleteSessionKey(c, "reset-password")
			restBeanRegister = models.SuccessRestBeanWithData("密码重置成功")
			models.UpdateUserStatus(requestData.Email, 0)
		} else {
			restBeanRegister = models.FailureRestBeanWithData(500, "内部错误，请联系管理员")
		}
	} else {
		restBeanRegister = models.FailureRestBeanWithData(http.StatusBadRequest, "请先完成邮箱认证")
	}

	//返回结果给前端
	c.JSON(restBeanRegister.Status, restBeanRegister)
}
