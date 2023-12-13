package models

import "3gnx/dao"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
	Xuehao   string `json:"xuehao"`
	Class    string `json:"class"`
}

/*
	Todo这个Model的增删改查操作都放在这里
*/
// CreateATodo 创建user
func CreateAUser(user *User) (err error) {
	err = dao.DB.Create(&user).Error
	return
}

func GetAllUser() (userList []*User, err error) {
	if err = dao.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func FindAUser(username string) (user *User, err error) {
	user = new(User)
	if err = dao.DB.Debug().Where("username=?", username).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *User) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&User{}).Error
	return
}
