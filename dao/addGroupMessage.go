package dao

import (
	"im/global"
	"im/model"

	"gorm.io/gorm"
)

func CreateAddGroupMessage(msg *model.AddGroupMessage) *gorm.DB {
	return global.DB.Create(msg)
}

func GetAddGroupMessagesByGroupId(id uint) ([]*model.AddGroupMessage, *gorm.DB) {
	list := make([]*model.AddGroupMessage, 0)
	db := global.DB.Where("group_id", id).Find(&list)
	return list, db
}

func DeleteAddGroupMessageBySenderIdAndGroupId(msg *model.AddGroupMessage) *gorm.DB {
	return global.DB.Where("sender_id=? and group_id=?", msg.SenderId, msg.GroupId).Delete(msg)
}

func GetAddGroupMessagesByGroupIds(ids []uint) ([]*model.AddGroupMessage, *gorm.DB) {
	ls := make([]*model.AddGroupMessage, 0)
	db := global.DB.Where("group_id in ?", ids).Find(&ls)
	return ls, db
}