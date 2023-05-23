package dao

import (
	"im/global"
	"im/model"
	"gorm.io/gorm"
)

func GetUserList() []*model.User {
	var list []*model.User
	global.DB.Find(&list)
	return list
}

func CreateUser(user *model.User) *gorm.DB {
	return global.DB.Create(user)
}

func DeleteUser(user *model.User) *gorm.DB {
	return global.DB.Delete(user)
}

func UpdateUser(user *model.User) *gorm.DB {
	return global.DB.Model(user).Updates(model.User{
		Name: user.Name,
		Introduce: user.Introduce,
	})
}

func GetUserByName(user *model.User) *gorm.DB {
	return global.DB.Where("name = ?", user.Name).First(user)
}

func GetUserById(user *model.User) *gorm.DB {
	return global.DB.Where("id = ?", user.ID).First(user)
}

func GetUserNameById(user *model.User) *gorm.DB {
	return global.DB.Select("name").Where("id=?", user.ID).First(user)
}

func GetUserIdByName(user *model.User) *gorm.DB {
	return global.DB.Select("id").Where("name=?", user.Name).First(user)
}

func GetFuzzyUserByUserName(name string) (users []*model.User, db *gorm.DB){
	db = global.DB.Where("name like ?", "%"+name+"%").Find(&users)
	return
}

func GetUsersByIds(ids []uint) (users []*model.User, db *gorm.DB){
	db = global.DB.Where("id in ?", ids).Find(&users)
	return
}