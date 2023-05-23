package dao

import (
	"im/global"
	"im/model"

	"gorm.io/gorm"
)

// 创建加好友的信息
func CreateAddUserMessage(msg *model.AddUserMessage) *gorm.DB {
	return global.DB.Create(msg)
}

// 根据接收方的id获取所有加好友的信息
func GetAddUserSenderIdsByReveiverId(id uint) ([]uint, *gorm.DB) {
	ids := make([]uint, 0)
	db := global.DB.Model(&model.AddUserMessage{}).Select("sender_id").Where("receiver_id = ?", id).Find(&ids)
	return ids, db
}

// 根据发送方的名称获取所有加好友的信息
func GetAddUserMessagesBySenderId(id uint) ([]*model.AddUserMessage, *gorm.DB) {
	var list []*model.AddUserMessage
	db := global.DB.Where("sender_id = ?", id).Find(list)
	return list, db
}

// 删除加好友的信息
func DeleteAddUserMessage(senderId, receiverId uint) *gorm.DB {
    return global.DB.Where("sender_id = ? or receiver_id = ?", senderId, receiverId).Delete(&model.AddUserMessage{})
}