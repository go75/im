package dao

import (
	"im/global"
	"im/model"

	"gorm.io/gorm"
)

func CreateUserSession(session *model.UserSession) *gorm.DB {
	return global.DB.Create(session)
}

func DeleteUserSession(UserSession *model.UserSession) *gorm.DB {
	return global.DB.Delete(UserSession)
}

func GetUserSessions(id uint) ([]*model.UserSession, *gorm.DB) {
	list := make([]*model.UserSession, 0)
	db := global.DB.Where("small_id = ? or big_id = ?", id, id).Find(&list)
	return list, db
}