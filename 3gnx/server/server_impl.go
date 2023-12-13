package server

import (
	"3gnx/dao"
	"3gnx/middles"
	"3gnx/models"
)

// 普通用户的注册
func UserRegister(user models.User, code string, sessionId string) string {
	//连接redis
	RedisClient, err := dao.ConnectToRedis()
	// 检查Redis中键是否存在
	key := "email" + sessionId + ":" + user.Email + ":false"
	exist, err := RedisClient.Exists(key).Result()
	if err != nil {
		return "内部错误，请联系管理员"
	}
	if exist == 0 {
		return "请先请求一封验证码邮件"
	}
	// 获取Redis中键对应的值
	result, err := RedisClient.Get(key).Result()
	if err != nil {
		return "内部错误，请联系管理员"
	}
	if result == "" {
		return "验证码失效，请重新请求"
	}
	if result == code {
		users, _ := models.FindAUser(user.Username)
		if users != nil {
			return "此用户名已被注册，请更换用户名"
		}
		RedisClient.Del(key)
		// 对密码进行加密
		privateencode := middles.Encode(user.Password)
		user.Password = privateencode
		// 创建新用户
		err := models.CreateAUser(&user)
		if err != nil {
			return "内部错误，请联系管理员"
		}

		return "" // 注册成功，返回空字符串表示成功
	} else {
		return "验证码错误，请检查后再提交"
	}
}

// 普通用户的登录
func UserLogin(username string, password string) string {
	//判断用户名是不是为空
	if username == "" {
		return "用户名不能为空"
	}
	// 根据用户名从数据库中获取用户信息
	user, _ := models.FindAUser(username)
	//if err != nil {
	//	return "内部错误，请联系管理员"
	//}

	// 验证用户是否存在
	if user == nil {
		return "用户不存在"
	}
	newpassword := middles.Encode(password)
	// 验证密码是否正确
	if newpassword != user.Password {
		return "密码不正确"
	}

	return "" // 登录成功，返回空字符串表示成功
}
