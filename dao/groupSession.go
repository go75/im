package dao

import (
	"im/global"
	"im/model"

	"gorm.io/gorm"
)

func CreateGroupSession(session *model.GroupSession) *gorm.DB {
	return global.DB.Create(session)
}

func DeleteGroupSession(GroupSession *model.GroupSession) *gorm.DB {
	return global.DB.Delete(GroupSession)
}

func QueryGroupSessionsByUserId(userId uint) ([]*model.GroupSession, *gorm.DB){
	res := make([]*model.GroupSession, 0)
	db := global.DB.Where("user_id", userId).Find(&res)
	return res, db
}

func QueryGroupIdsByUserId(userId uint) ([]uint, *gorm.DB) {
	res := make([]uint, 0)
	db := global.DB.Model(&model.GroupSession{}).Select("group_id").Where("user_id=?", userId).Find(&res)
	return res, db
}

func QueryGroupSessionsByGroupId(session *model.GroupSession) ([]*model.GroupSession, *gorm.DB) {
	res := make([]*model.GroupSession, 0)
	db := global.DB.Where("group_id", session.GroupId).Find(&res)
	return res, db	
}

func QueryGroupSessionUserIdsByGroupId(session *model.GroupSession) ([]*model.GroupSession, *gorm.DB){
	res := make([]*model.GroupSession, 0)
	db := global.DB.Select("user_id").Where("group_id=?", session.GroupId).Find(&res)
	return res, db
}