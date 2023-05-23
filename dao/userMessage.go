package dao

import (
	"im/global"
	"im/log"
	"im/model"

	"gorm.io/gorm"
)

func CreateUserMessage(msg *model.UserMessage) *gorm.DB {
	return global.DB.Create(msg)
}

func QueryUserMessageByIds(id1 uint, id2 uint) ([]*model.UserMessage, *gorm.DB) {
	var list []*model.UserMessage
	db := global.DB.Where("(sender_id=? and receiver_id=?) or (receiver_id=? and sender_id=?)", id1, id2, id1, id2).Order("id").Find(&list)
	for _, l := range list {
		log.Info.Println(*l)
	}
	return list, db
}