package dao

import (
	"im/global"
	"im/model"

	"gorm.io/gorm"
)

// 创建群聊消息
func CreateGroupMessage(msg *model.GroupMessage) *gorm.DB {
	return global.DB.Create(msg)
}

// 查询群聊所有消息
func QueryGroupMessage(groupId uint) ([]*model.GroupMessage, *gorm.DB) {
	var list []*model.GroupMessage
	db := global.DB.Where("group_id=?", groupId).Order("id").Find(&list)
	return  list, db
}