package controller

import (
	"3gnx/middles"
	"3gnx/models"
	"3gnx/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户注册并返回信息给客户端
func UserRegister(c *gin.Context) {

	sessionID := middles.GetSessionId(c.Writer, c.Request)
	var requestData struct {
		User models.User `json:"user"`
		Code string      `json:"code"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用UserRegister函数进行注册
	result := server.UserRegister(requestData.User, requestData.Code, sessionID)

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
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
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
